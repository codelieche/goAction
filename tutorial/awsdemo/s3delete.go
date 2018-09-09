package awsdemo

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

/**
从aws s3中删除对象
*/
func DeleteS3Object(bucketName, keyName string) (bool, error) {
	// 1. 先实例化对象
	sess := NewAwsSession()
	svc := s3.New(sess)

	// 2. 删除对象
	input := s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
	}
	if _, err := svc.DeleteObject(&input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
