package monitor

import "log"

// 系统信息统计的channel
const (
	taskCount          = 100
	taskExecuteSuccess = 1
	taskExecuteError   = 0
	taskEventCreate    = 10
)

var systemInfoStatChan = make(chan int, 100)

func (process *Process) statSystemInfo() {
	for n := range systemInfoStatChan {
		switch n {
		case taskCount:
			process.Info.Count += 1

		case taskExecuteSuccess:
			process.Info.SuccessNum += 1

		case taskExecuteError:
			process.Info.ErrorNum += 1

		case taskEventCreate:
			process.Info.EventCount += 1

		default:
			log.Println("systemInfoStatChan：传入的值有误！", n)
		}

	}
}
