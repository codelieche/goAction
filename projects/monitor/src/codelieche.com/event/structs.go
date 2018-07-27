package event

import "time"

type Event struct {
	Monitor    int       `json:"monitor"` // 监控的ID
	Title      string    `json:"title"`   // 事件的标题
	Content    string    `json:"content"` // 事件内容
	Level      int       `json:"level"`   // 事件级别，级别越高越严重
	Status     string    `json:"status"`  // 状态：toto, doing, autofixed
	StatusCode int       // 状态码，数字，方便排序
	Creator    string    //事件的创建者
	Handler    string    // 事件的处理者
	TimeStart  time.Time // 事件开始时间
}
