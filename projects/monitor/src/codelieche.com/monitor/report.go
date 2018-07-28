package monitor

import (
	"encoding/json"

	"codelieche.com/settings"
	"github.com/levigross/grequests"
)

/**
报告事件
*/

// 报告监控事件去web接口

func (web *HandleWebEvent) Report(e *Event) (bool, string) {
	// 向后端报告监控事件
	var result bool

	// 第1步：登录系统并实例化 session
	session, err := settings.Login()

	if err != nil {
		return false, err.Error()
	}

	// 第2步：添加Evnent
	// 2-1: 把event的数据变成json的
	if e.Status == "" {
		e.Status = "todo"
	}
	data, err := json.Marshal(e)
	if err != nil {
		return false, err.Error()
	}

	// 2-2：构造创建evnet的请求选项
	ro := &grequests.RequestOptions{
		JSON:    data,
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	// 2-3: 发起请求
	resp, err := session.Post(settings.Config.Web.EventCreateUrl, ro)

	if err != nil {
		return false, err.Error()
	}

	// 第3步：判断响应的结果是否有id字段
	eventCreated := Event{}
	// 把字符串转换成 Event struct
	if err := json.Unmarshal(resp.Bytes(), &eventCreated); err != nil {
		return false, err.Error()
	} else {
		// 判断创建的事件，是否正确
		if eventCreated.Monitor == e.Monitor {
			result = true
			// 只有执行到这里：才表示创建event成功
			return result, resp.String()
		} else {
			// 创建事件失败，有问题
			return false, resp.String()
		}
	}
}
