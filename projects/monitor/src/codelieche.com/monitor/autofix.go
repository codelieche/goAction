package monitor

import (
	"encoding/json"

	"codelieche.com/settings"
	"github.com/levigross/grequests"
)

/**
当监控由异常变成正常的时候，需要自动修复下
*/

/**
自动修复web监控
1. 登录web系统
2. 取出web系统当前监控相关的所有正在处理中的事件
3. 设置其状态为autofixed
*/
func (web *HandleWebEvent) AutoFix(m *Monitor) (FixResponse, error) {
	// 第1步：先准备好接收响应的接口
	result := FixResponse{}

	// 第2步：登录系统获取session
	session, err := settings.Login()
	if err != nil {
		return result, err
	}

	// 第3步：发起Post请求，autofix
	// 3-1: 构造请求选项
	ro := &grequests.RequestOptions{
		JSON:    map[string]int{"monitor": m.Id},
		Headers: map[string]string{"Content-Type": "application/json"},
	}
	// 3-2: 发起请求
	if resp, err := session.Post(settings.Config.Web.EventAutoFixUrl, ro); err != nil {
		return result, err
	} else {
		// 3-3: 响应成功了，构造响应数据：json to struct
		if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
			return result, err
		} else {
			// 只有执行到这里才表示成功修复了监控事件
			return result, nil
		}
	}
	return result, nil
}
