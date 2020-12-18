package awsdemo
/**
aws发送消息和删除消息
 */
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

/**
获取消息队列的Url
*/

func GetAwsQueueUrl(queueName string) (string, error) {
	// 1. 先实例化
	sess := NewAwsSession()
	svc := sqs.New(sess)

	// 2. 获取消息队列的地址
	var queueUrl string
	input := sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}
	if result, err := svc.GetQueueUrl(&input); err != nil {
		return queueUrl, err
	} else {
		queueUrl = *result.QueueUrl
		return queueUrl, nil
	}
}

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


/**
获取消息
*/
func ReceiveMessage(queueUrl string, maxNumber int64) (result *sqs.ReceiveMessageOutput, err error) {
	// 获取消息

	// 1. 先实例化
	sess := NewAwsSession()
	svc := sqs.New(sess)

	// 2. 接收消息
	receiveInput := sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(maxNumber),
		VisibilityTimeout:   aws.Int64(20),
		WaitTimeSeconds:     aws.Int64(0),
	}

	result, err = svc.ReceiveMessage(&receiveInput)

	// 3. 对结果进行校验
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// 从sqs中删除消息
func DeleteSqsMessage(queueUrl string, receiptHandle string) error {
	// 1. 先实例化
	sess := NewAwsSession()
	svc := sqs.New(sess)

	// 2. 操作删除
	input := sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: aws.String(receiptHandle),
	}

	if _, err := svc.DeleteMessage(&input); err != nil {
		return err
	} else {
		return nil
	}
}

