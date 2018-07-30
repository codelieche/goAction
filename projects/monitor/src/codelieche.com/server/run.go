package server

import (
	"log"
	"sync"

	"codelieche.com/event"
	"codelieche.com/execute"
	"codelieche.com/logs"
	"codelieche.com/monitor"
	"codelieche.com/source"
	"time"
)

/**
服务启动入口
*/

func Run() {
	log.Println("程序开始运行")

	// 获取列表的Source:
	// 需要实现monitor.Lister的接口
	web := source.ListMonitorFromWeb{}

	// 执行监控任务的执行器：
	// 需要实现monitor.Executer的接口
	webExecute := execute.WebTaskExecute{}

	// 执行结果信息 映射
	executeInfoMapData := make(map[int]monitor.ExecuteInfo)

	executeInfoMap := monitor.ExecuteInfoMap{
		Data: &executeInfoMapData,
		Lock: &sync.RWMutex{},
	}

	// 执行任务的channel
	taskChan := make(chan (monitor.Task), 10)
	logChan := make(chan (monitor.Log), 10)

	// 处理异常事件的handler
	eventHandler := event.HandleWebEvent{}

	// 处理日志的handler
	logHandle := logs.RecordLogToInfluxdb{}

	// 监控程序的系统信息
	processInfo := monitor.SystemInfo{
		StartTime: time.Now(),
	}
	process := monitor.Process{
		Source:         &web,
		TaskExecute:    &webExecute,
		ExecuteInfoMap: &executeInfoMap,
		TaskChan:       taskChan,
		LogChan:        logChan,
		EventHandle:    &eventHandler,
		LogHandle:      &logHandle,
		Info:           &processInfo,
	}

	// 执行process 程序
	go process.Run()

	// 实例化webServer
	monitorServer := MonitorServer{
		startTime: time.Now(),
		process:   &process,
	}
	monitorServer.start()
}
