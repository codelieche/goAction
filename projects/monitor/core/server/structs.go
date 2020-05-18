package server

import (
	"goAction/projects/monitor/core/monitor"
	"time"
)

type MonitorServer struct {
	startTime time.Time
	process   *monitor.Process
}
