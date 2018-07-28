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
