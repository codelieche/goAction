package awsdemo

import "testing"

/**
发送消息到aws测试
*/

func TestSendMessageToSQS(t *testing.T) {
	queueName := "codelieche_message"
	body := `{"status": true, "message": "成功发送消息到aws"}`

	if err := SendMessageToSQS(queueName, body); err != nil {
		t.Error("发送消息出错：", err.Error())
	} else {
		t.Log("发送消息到aws成功！")
	}
}
