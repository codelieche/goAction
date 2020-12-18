package awsdemo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Config = ProjectConfig{
	AwsRegion: "ap-southeast-1",
}

/**
获取aws的session
*/
func NewAwsSession() *session.Session {
	// 1. 先获取到需要的key
	aws_secret_key_id := Config.AwsSecretKeyId
	aws_secret_access_key := Config.AwsSecretAccessKey

	// 2. 实例化config
	// 如果key为空 用系统中配置的key
	var config aws.Config
	if aws_secret_key_id == "" || aws_secret_access_key == "" {
		config = aws.Config{
			Region: aws.String(Config.AwsRegion),
		}
	} else {
		config = aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(aws_secret_key_id, aws_secret_access_key, ""),
		}
	}

	// 3. 实例化session
	sess := session.Must(session.NewSession(&config))
	return sess
}
