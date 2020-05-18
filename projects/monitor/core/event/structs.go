package event

import "goAction/projects/monitor/core/monitor"

// 需要实现monitor的Reporter和AutoFixer接口
// 处理Web事件的操作
type HandleWebEvent struct {
	// 其要实现：Handler的接口
	// 那么需要实现两个方法：Report, AutoFix
	monitor.Handler
}
