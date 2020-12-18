package awsdemo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
)

/**
从aws的bucket中下载图片
*/

func GetObjectFromS3(bucketName, keyName string) (*aws.WriteAtBuffer, error) {
	// 第1步：实例化session
	sess := NewAwsSession()

	// 第2步：实例化downloader
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Write the contents of Object to the file
	buff := &aws.WriteAtBuffer{}

	fileObject := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
	}

	// 下载
	if _, err := downloader.Download(buff, fileObject); err != nil {
		return nil, fmt.Errorf("下载文件出错, %v", err)
	} else {
		return buff, nil
	}
}

