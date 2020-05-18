package logs

import "goAction/projects/monitor/core/monitor"

// 记录日志到influxdb中
// 需要实现：monitor.LogHandler的接口
// 也就是要实现：RecordLog(c chan Log)方法
type RecordLogToInfluxdb struct {
	monitor.LogHandler
}

// 日志直接输出到标准输出中
type RecordLogToStdOut struct {
	monitor.LogHandler
}
