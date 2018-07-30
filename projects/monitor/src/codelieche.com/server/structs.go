package server

import (
	"codelieche.com/monitor"
	"time"
)

type MonitorServer struct {
	startTime time.Time
	process   *monitor.Process
}
