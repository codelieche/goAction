package server

import (
	"log"

	"time"

	"codelieche.com/monitor"
)

/**
服务启动入口
*/

func Run() {
	log.Println("程序开始运行")
	// 登录
	//sesession, err := settings.Login()
	//log.Println(sesession, err)
	// 获取列表
	web := monitor.ListMonitorFromWeb{}

	process := monitor.Process{
		Source: &web,
	}
	monitors, err := process.Source.List()
	if err != nil {
		log.Println("获取监控数据列表出错：", err.Error())
	}
	execute := monitor.WebTaskExecute{}

	for _, monitor_i := range monitors {
		//log.Println(monitor_i)
		task := monitor.Task{
			Monitor:      &monitor_i,
			Status:       "todo",
			ExecutedTime: time.Now().Add(time.Duration(3) * time.Second),
			ExpiredTime:  time.Now().Add(time.Duration(20) * time.Second),
		}
		execute.Execute(&task)
		break
	}
}
