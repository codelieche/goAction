package monitor

import (
	"log"
	"time"
)

/**
执行监控任务
1. 不断的从监控任务的channel中取出数据
2. 执行监控任务
*/
func (process *Process) ExecuteMonitorTask() {
	for {
		//var task Task
		//task <- process.TaskChan
		task := <-process.TaskChan
		go process.executeTask(&task)
	}
}

/**
记录监控日志
1. 不断的从日志的channel中取出数据
2. 写入Log
*/
func (process *Process) RecordLog() {
	//for {
	//	executeLog := <-process.LogChan
	//	log.Println("处理日志：", executeLog)
	//}

	// 使用传入的：LogHandle.RecordLog()处理
	process.LogHandle.RecordLog(process.LogChan)
}

/**
处理程序：生成监控任务
1. 先从Source.List()中取出监控的列表
2. 根据监控列表生成Task
3. 把Task加入到process.TaskChan中
*/
func (process *Process) generateTaskMain(monitorListCached *[]Monitor, nextFreshTime time.Time) {
	// 第1步：先从Source.List()中取出监控列表数据
	if monitorList, err := process.Source.List(); err != nil {
		// 获取列表数据出错
		log.Println("获取监控列表数据出错:", err.Error())
		// 使用缓存：
		for _, m := range *monitorListCached {
			// 生成任务
			process.generateMonitorTask(m, nextFreshTime)
		}
	} else {
		// 修改缓存
		*monitorListCached = monitorList
		for _, m := range monitorList {
			// 生成任务
			process.generateMonitorTask(m, nextFreshTime)
		}
	}
}

/**
根据Monitor -> Task
*/
func (process *Process) generateMonitorTask(m Monitor, nextFreshTime time.Time) {
	// 第1步：判断监控是否启用
	if !m.IsActive {
		return
	}
	// 第2步：下面开始生成任务去process.TaskChan中
	// 2-1: 取出监控的执行信息
	monitorId := m.Id
	var executeInfo ExecuteInfo
	now := time.Now()
	executeInfo = process.ExecuteInfoMap.Get(monitorId)
	if executeInfo.ExecuteTime.Year() < 2000 {
		//log.Println("第一次处理这个monitor")
		// 设置执行信息
		executeInfo = ExecuteInfo{
			IsOk:            true,
			ExecuteTime:     now,
			LastExecuteTime: now,
			ExecuteCount:    0,
			ErrorCount:      0,
		}
		// 设置执行信息到map中
		process.ExecuteInfoMap.Set(monitorId, executeInfo)
	}

	// 如果执行信息中的：最后执行时间小于now了，就把LastTaskTime设置为now
	// 这个是为了防止程序运行中，待机了，进入休眠停止很久了，再次继续执行，而LastTaskTime还是很久之前的
	if executeInfo.LastExecuteTime.Before(now) {
		// 这个监控最后一次执行监控任务的事件
		executeInfo.LastExecuteTime = now
	}

	// 2-2: 对上次执行任务的时间进行修复处理
	// 任务执行时间、任务执行过期时间
	taskExecuteTime := executeInfo.LastExecuteTime

	// 任务执行过期时间
	taskExecuteExpireTime := taskExecuteTime.Add(time.Duration(m.Interval) * time.Second)

	// 第3步：创建Task
	// 注意：创建的条件是：任务执行时间在下次刷新时间的前面
	for taskExecuteTime.Before(nextFreshTime) {
		// 3-1：重新设置任务执行时间、过期时间
		taskExecuteTime = executeInfo.LastExecuteTime.Add(time.Duration(m.Interval) * time.Second)
		taskExecuteExpireTime = executeInfo.LastExecuteTime.Add(time.Duration(m.Interval*2) * time.Second)
		// 3-2: 创建监控任务
		task := Task{
			Monitor:      &m,
			Status:       "todo",
			ExecutedTime: taskExecuteTime,
			ExpiredTime:  taskExecuteExpireTime,
		}
		// log.Println(taskExecuteTime, executeInfo.LastExecuteTime)

		// 修改执行信息的最后一次任务时间
		executeInfo.LastExecuteTime = taskExecuteTime
		// 把执行信息更新到 map中
		process.ExecuteInfoMap.Set(monitorId, executeInfo)

		// 第4步：把任务加入到channel中
		process.TaskChan <- task
		// 加入条数据到：统计信息的channel中
		systemInfoStatChan <- taskCount
	}
}
