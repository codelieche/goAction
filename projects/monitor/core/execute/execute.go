package execute

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"regexp"
	"strings"

	"goAction/projects/monitor/core/monitor"

	"github.com/levigross/grequests"
)

/**
执行监控相关任务
*/

func (execute *WebTaskExecute) Execute(task *monitor.Task) (*monitor.Result, error) {
	// 执行监控任务
	//log.Println("执行监控任务：", task.Monitor.Id, task.Monitor.Name)

	// 先准备结果
	result := monitor.Result{
		Success:  false,
		Executed: false,
	}

	// 第1步：先对执行的监控类型和执行时间做校验
	// 1-1：执行任务的监控类型如果不是web就取消
	if task.Monitor.Category != "web" {
		return nil, fmt.Errorf("需要执行web相关的监控才行：当前的监控类型为:%s", task.Monitor.Category)
	}
	// 1-2: 如果任务过期了，也无需再执行
	if task.ExpiredTime.Before(time.Now()) {
		return nil, nil
	}

	// 第2步：判断执行时间
	// 如果执行时间在当前时间之后，需要延时一下
	if task.ExpiredTime.After(time.Now()) {
		time.Sleep(task.ExecutedTime.Sub(time.Now()))
	}

	// 第3步：开始执行任务
	// 3-1: 实例化session
	sRo := &grequests.RequestOptions{
		//RequestTimeout: time.Duration(5000 * time.Millisecond),
	}
	session := grequests.NewSession(sRo)

	// 3-2: 去执行每个步骤
	var timeCount float64
	for i, step := range task.Monitor.Steps {
		// 第4步：执行操作
		//log.Printf("执行步骤: %d\n", i+1)
		// 4-1: 判断是否需要延时
		if step.Sleep > 0 {
			// 睡眠时间是毫秒
			time.Sleep(time.Duration(step.Sleep) * time.Millisecond)
		}
		// 4-2: 访问
		timeStart := time.Now() // 记录下开始时间，计算本次请求的事件要用到
		ro := &grequests.RequestOptions{
			RequestTimeout: time.Duration(step.Timeout) * time.Millisecond,
			Headers:        map[string]string{},
		}
		var resp *grequests.Response
		var err error
		switch step.Method {
		case "GET":
			resp, err = session.Get(step.Url, ro)
		case "POST":
			// 如果是post方法，就需要处理下请求数据
			(*ro).RequestBody = bytes.NewReader([]byte(step.Data))
			(*ro).Headers["Content-Type"] = "application/x-www-form-urlencoded"
			resp, err = session.Post(step.Url, ro)
		default:
			log.Println("暂时不支持这个请求的方法")
			return nil, nil
		}

		// 4-3: 判断响应是否有误
		if !result.Executed {
			result.Executed = true
		}
		// 增加下执行请求时间
		timeCount += float64(time.Since(timeStart).Nanoseconds())
		if err != nil {
			log.Println("执行出现了错误:", err)
			// 构造事件
			event := monitor.Event{
				Monitor: task.Monitor.Id,
				Title:   fmt.Sprintf("第%d步：执行出现错误", i+1),
				Content: err.Error(),
				Level:   task.Monitor.Level, // 时间级别跟监控配置的级别一样
			}
			result.Event = &event
			result.Elapsed = timeCount / 1000000
			return &result, err
		}

		// 第5步：对请求进行校验
		// 5-1: 开始对响应的状态码进行校验：检查结果是bool型
		if !CheckResponseStatusCode(&step, resp.StatusCode) {
			log.Println("检查状态码结果为False")
			//log.Println(step)
			message := fmt.Sprintf("请求:%s, 响应码是%d, 检查条件是:%s %d 结果为false",
				step.Url, resp.StatusCode, step.CodeMinExpr, step.CodeMin)
			//log.Println(message)
			if step.CodeMax > 0 {
				message = fmt.Sprintf("请求:%s，响应码是%d，检查条件是：%s %d, %s %d 结果为false",
					step.Url, resp.StatusCode, step.CodeMinExpr, step.CodeMin, step.CodeMaxExpr, step.CodeMax)
			}
			// 构造事件
			event := monitor.Event{
				Monitor: task.Monitor.Id,
				Title:   fmt.Sprintf("状态码异常(第%d步)", i+1),
				Content: message,
				Level:   task.Monitor.Level,
			}
			result.Event = &event
			result.Elapsed = timeCount / 1000000
			return &result, nil
		} else {
			//log.Println("检查状态码成功")
		}

		// 5-2: 检查标题：不为空才检查
		if step.Title != "" {
			reTitle, _ := regexp.Compile(`<title>.*</title>`)
			title := reTitle.FindString(resp.String())
			// 正则取到title然后判断是否包含标题
			if !strings.Contains(title, step.Title) {
				// 响应的标题中没有，步骤中的Title
				// 构造事件
				message := fmt.Sprintf("请求:%s, 响应的标题中不包含：%s,返回的标题是: %s",
					step.Url, step.Title, title)
				event := monitor.Event{
					Monitor: task.Monitor.Id,
					Title:   fmt.Sprintf("检查响应标题(第%d步)出错", i+1),
					Content: message,
					Level:   1,
				}
				result.Event = &event
				result.Elapsed = timeCount / 1000000
				return &result, nil
			} else {
				//fmt.Println("检查标题成功")
			}
		}

		// 5-3: 检查主体内容
		if step.Body != "" && !strings.Contains(resp.String(), step.Body) {
			// 响应的内容中没有，步骤中的内容
			// 构造事件
			message := fmt.Sprintf("请求:%s, 响应的内容中不包含：%s", step.Url, step.Body)
			event := monitor.Event{
				Monitor: task.Monitor.Id,
				Title:   fmt.Sprintf("检查响应内容(第%d步)", i+1),
				Content: message,
				Level:   1,
			}
			result.Event = &event
			result.Elapsed = timeCount / 1000000
			return &result, nil
		}
	}

	//  第6步：返回结果
	// 本次请求总共的事件
	timeCount = timeCount / 1000000
	//log.Println("本次请求总共用了", timeCount, "毫秒")
	result.Elapsed = timeCount
	// 执行到这表示成功了
	result.Success = true
	return &result, nil
}
