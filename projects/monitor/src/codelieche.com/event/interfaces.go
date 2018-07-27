package event

/**
事件相关的接口
1. 创建事件去web
2. 也可以直接发送消息到phone等
3. 不同的报告方式，实现了Reporter接口即可
*/

type Reporter interface {
	Report() (bool, string)
}
