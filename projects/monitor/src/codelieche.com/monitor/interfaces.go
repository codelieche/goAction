package monitor

/*
监控相关的接口
*/

// 获取监控的列表
type Lister interface {
	List() ([]Monitor, error)
}

/*
有个这个Lister:
- 可以从web后台获取监控的列表
- 可以从etcd、redis、数据库中获取监控列表
- 可以从配置文件中读取监控列表
- 这样增加了程序的可拓展性
*/

// 执行监控的接口
type Executer interface {
	Execute(task *Task) (*Result, error)
}

// 报告事件
type Reporter interface {
	Report(event *Event) (bool, string)
}

/**
事件相关的接口
1. 报告监控事件
2. 修复监控
*/

// 设置监控已经恢复
type AutoFixer interface {
	AutoFix(m *Monitor) (FixResponse, error)
}

// 事件相关的处理接口
type Handler interface {
	Reporter
	AutoFixer
}

// 日志操作的接口
type LogHandler interface {
	RecordLog(c chan Log)
}
