package monitor

import (
	"log"

	"encoding/json"
	"fmt"

	"codelieche.com/settings"
)

// 获取监控的列表

/**
从web后端：获取监控列表
- 通过运维平台获取监控列表
1. 先登录web系统
2. 获取所有的监控列表数据：list是分页的，all是返回所有的监控列表
*/
type ListMonitorFromWeb struct {
	//Session *grequests.Session // 登录了web系统的session，然后可以直接访问页面获取list
	//Url     string
}

func (web *ListMonitorFromWeb) List() ([]Monitor, error) {
	// 获取监控列表
	// 发起请求：得到所有的监控列表

	// 第1步：登录系统并实例化session
	session, err := settings.Login()
	if err != nil {
		return nil, err
	}

	// 第2步：获取监控数据，获取列表
	if resp, err := session.Get(settings.Config.Web.GetMonitorListUrl(), nil); err != nil {
		// 出现错误
		log.Println("获取监控列表的时候出现错误")
		log.Println(err.Error())
		return nil, err
	} else {
		// 第3步：处理返回的响应
		// 后台使用的是：Django Rest Framework
		var results []Monitor
		if err := json.Unmarshal(resp.Bytes(), &results); err == nil {
			// 第4步：返回监控列表【只有执行到这里才是正确的
			if len(results) == 0 {
				log.Println("获取监控列表为空([])")
			}
			return results, nil
		} else {
			return nil, fmt.Errorf("解析列表出错:%s", err)
		}
	}
	return nil, fmt.Errorf("执行到最后，出错了")
}

/**
从etcd中获取监控列表
*/
