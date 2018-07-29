package source

import "codelieche.com/monitor"

/**
从web后端：获取监控列表
- 通过运维平台获取监控列表
1. 先登录web系统
2. 获取所有的监控列表数据：list是分页的，all是返回所有的监控列表
*/

// 列出监控列表：需要实现Lister的接口
type ListMonitorFromWeb struct {
	//Session *grequests.Session // 登录了web系统的session，然后可以直接访问页面获取list
	//Url     string
	monitor.Lister
}
