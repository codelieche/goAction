package server

import (
	"log"

	"codelieche.com/monitor"
)

/**
服务启动入口
*/

func Run() {
	log.Println("程序开始运行")
	// 获取列表的Source
	web := monitor.ListMonitorFromWeb{}
	// 执行监控任务的执行器
	execute := monitor.WebTaskExecute{}
	// 执行结果信息 映射
	executeInfoMap := monitor.ExecuteInfoMap{}
	// 执行任务的channel
	taskChan := make(chan (monitor.Task), 10)
	logChan := make(chan (monitor.Log), 10)

	process := monitor.Process{
		Source:         &web,
		TaskExecute:    &execute,
		ExecuteInfoMap: &executeInfoMap,
		TaskChan:       taskChan,
		LogChan:        logChan,
	}

	// 执行process 程序
	process.Run()
}
