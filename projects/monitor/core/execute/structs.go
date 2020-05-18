package execute

import "goAction/projects/monitor/core/monitor"

// web监控任务执行器:
// 需要实现：monitor.Executer的接口
type WebTaskExecute struct {
	monitor.Executer
}
