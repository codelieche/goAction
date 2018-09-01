package awsdemo
/**
aws发送消息和删除消息
 */
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

/**
发送消息到消息队列中
*/

// 发送消息到SQS
// 参数：body 传递的消息内容，queueName 消息队列名称
func SendMessageToSQS(queueName string, body string) error {
	// 1. 先实例化session
	sess := NewAwsSession()

	svc := sqs.New(sess)

	// 2. 获取消息队列地址
	var queueUrl string

	if result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}); err != nil {
		// 获取消息队列地址出错
		return err
	} else {
		queueUrl = *result.QueueUrl
	}

	// 3. 发送消息
	msgInput := sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(queueName),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("codelieche"),
			},
		},
		MessageBody: aws.String(body),
		QueueUrl:    &queueUrl,
	}
	if _, err := svc.SendMessage(&msgInput); err != nil {
		return err
	} else {
		// 发送消息成功，返回nil
		return nil
	}

}