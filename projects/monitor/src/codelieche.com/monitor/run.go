package monitor

import (
	"log"
	"time"

	"codelieche.com/settings"
)

/**
监控 Process 的执行入口
*/

func (process *Process) Run() {
	// 监控处理结构体执行入库 函数
	log.Println("程序开始执行:", time.Now())

	// 启动执行任务的goroutine
	go process.ExecuteMonitorTask()
	// 启动处理执行日志的goroutine
	go process.RecordLog()

	// 启动系统统计信息
	go process.statSystemInfo()

	// 启动计算tps的函数
	//process.calculateTps()

	monitorCached := []Monitor{}
	for {
		start := time.Now()
		nextTime := start.Add(time.Duration(settings.Config.Web.Interval) * time.Second)
		// 执行生成task的程序
		process.generateTaskMain(&monitorCached, nextTime)
		time.Sleep(nextTime.Sub(time.Now()))
	}
}
