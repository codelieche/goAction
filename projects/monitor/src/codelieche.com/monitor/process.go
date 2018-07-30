package monitor

import (
	"fmt"
	"log"
	"time"
)

/**
处理程序执行监控任务
*/
func (process *Process) executeTask(task *Task) {
	// 第1步：先判断这个任务是否过期
	if task.ExpiredTime.Before(time.Now()) {
		log.Println("任务过期时间，已经过了，无需执行")
		return
	}
	// 第2步：如果执行时间在当前时间之后，需要先延时下
	if task.ExecutedTime.After(time.Now()) {
		time.Sleep(task.ExecutedTime.Sub(time.Now()))
	}

	// 第3步：开始执行任务：使用process的TaskExecute
	result, err := process.TaskExecute.Execute(task)
	if result == nil {
		log.Println("执行任务结果返回空:", task.Monitor)
		return
	}

	// 第4步：执行完毕，做相应的处理
	// 4-1: 判断err是否为空
	if err != nil {
		log.Println(err.Error())
	}

	// 4-2: 得到执行信息
	monitorId := task.Monitor.Id
	executeInfo := process.ExecuteInfoMap.Get(monitorId)
	// 4-3：对执行次数处理
	if result != nil && result.Executed {
		executeInfo.ExecuteCount += 1
	}
	// 把执行结果返回的Log加入到执行日志的chan中
	// 记录日志
	executeLog := Log{
		Id:       task.Monitor.Id,
		Time:     time.Now(),
		Success:  result.Success,
		Elapesed: result.Elapsed,
	}
	// 把log加入到channel中
	process.LogChan <- executeLog

	// 4-4: 判断是否执行成功
	if result.Success {
		// 执行成功：插入条数据到统计channel中
		systemInfoStatChan <- taskExecuteSuccess

		// 执行成功
		if !executeInfo.IsOk {
			executeInfo.IsOk = true
		}
		if executeInfo.ErrorCount > 0 {
			// 表示以前出现了异常，但是现在ok了

			// 如果执行次数小于Retries的次数，那么给其设置为0
			if executeInfo.ErrorCount <= task.Monitor.Retries {
				// 在重试次数内就已经正常了，直接给其错误计数设置为0
				executeInfo.ErrorCount = 0
			} else {
				// 表示大于重试次数后才恢复正常的，表示有异常事件，需要告警说问题恢复了
				log.Println(time.Now(), "服务恢复了")
				// 雍余了点，但是先这样，后续优化
				executeInfo.ErrorCount = 0

				// 执行监控异常事件的：AutoFix操作
				if result, err := process.EventHandle.AutoFix(task.Monitor); err != nil {
					log.Println("执行自动修复出错:", err.Error())
				} else {
					if result.Status {
						log.Println("执行自动修复成功：", result.Message)
					} else {
						log.Println("执行自动修复失败：", result.Message)
					}
				}
			}
		}
	} else {
		// 执行失败
		// 执行失败：插入条数据到统计channel中
		systemInfoStatChan <- taskExecuteError

		executeInfo.ErrorCount = executeInfo.ErrorCount + 1
		if executeInfo.IsOk {
			executeInfo.IsOk = false
		}
		if result.Executed {
			// 先判断是否执行了，这样的话就需要创建错误事件了
			// 第5步：处理异常

			// 5-1: 判断是不是第一次出现异常
			// 当错误几次，刚好等于重试次数的时候，就表示需要告警了
			// 因为把 错误统计次数放到了前面，所以这里是等于重试次数 + 1的时候就要告警了
			if executeInfo.ErrorCount == task.Monitor.Retries+1 {
				// 第6步：创建新的异常事件
				message := fmt.Sprintf("%s(%d): 需要创建异常事件了", task.Monitor.Name, task.Monitor.Id)
				// 执行失败：创建一条异常事件
				systemInfoStatChan <- taskEventCreate

				log.Println(message)
				// 创建异常事件
				// 执行监控返回的结果中有个Event对象的
				if success, message := process.EventHandle.Report(result.Event); success {
					// 创建异常事件成功
					msg := fmt.Sprintf("%s(%d): 创建异常事件成功!", task.Monitor.Name, task.Monitor.Id)
					log.Println(msg)
				} else {
					// 创建异常事件失败
					log.Println("创建异常事件失败：", message)
				}

			}
		} else {
			// 未执行，那么就无需处理
			return
		}
	}
	// 由于 := 得到的executeInfo是指传递，所以需要重新修改下
	// 最后重新设置执行信息
	process.ExecuteInfoMap.Set(monitorId, executeInfo)
}
