package awsdemo

import "testing"


// 测试获取消息队列的url
func TestGetAwsQueueUrl(t *testing.T) {
	var ququeName = "codelieche_message"

	if url, err := GetAwsQueueUrl(ququeName); err != nil {
		t.Error("获取sqs的Url失败：", err.Error())
	} else {
		t.Log("获取sqs的Url成功：", url)
	}
}

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

// 测试接收消息
func TestReceiveMessage(t *testing.T) {

	// 1. 先获取到消息队列的url
	var queueUrl string
	var ququeName = "codelieche_message"

	if url, err := GetAwsQueueUrl(ququeName); err != nil {
		t.Error("获取sqs的Url失败：", err.Error())
		return
	} else {
		t.Log("获取sqs的Url成功：", url)
		queueUrl = url
	}

	// 2. 接收消息
	if result, err := ReceiveMessage(queueUrl, 10); err != nil {
		t.Error("获取消息失败:", err.Error())
	} else {
		t.Log("获取消息成功，长度为：", len(result.Messages))
		for _, m := range result.Messages {
			t.Log(*(m.Body))
		}
	}
}

// 测试删除消息
func TestDeleteSqsMessage(t *testing.T) {
	queueUrl := "https://sqs.ap-southeast-1.amazonaws.com/xxxxx"
	receiptHandle := "xxxx"
	// receiptHandle 和 queueUrl可以从ReciveMessage中获取到

	if err := DeleteSqsMessage(queueUrl, receiptHandle); err != nil {
		t.Error("删除消息失败：", err.Error())
	} else {
		t.Log("删除消息成功！")
	}
}