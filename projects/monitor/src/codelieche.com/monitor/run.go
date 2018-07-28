package monitor

import (
	"log"
	"time"
)

/**
监控 Process 的执行入口
*/

func (process *Process) Run() {
	// 监控处理结构体执行入库 函数
	log.Println("程序开始执行:", time.Now())
}
