package awsdemo

import (
	"io"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"strings"
)

/**
上传图片/ 文件到bucket中
*/

func UpljoadBinFileToS3(f io.Reader, bucketName string, keyName, contentType string) (string, error) {
	// 第1步：获取session
	sess := NewAwsSession()

	// 第2步：实例化一个上传器
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// 第3步：开始上传
	if contentType == "" {
		contentType = "image/png"
	}

	inputObject := s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(keyName),
		Body:        f,
		ContentType: aws.String(contentType),
	}
	if result, err := uploader.Upload(&inputObject); err != nil {
		return "", fmt.Errorf("上传文件失败：%v", err)
	} else {
		// 得到对象的url
		url := strings.Split(result.Location, "amazonaws.com/")[1]
		return url, nil
	}
}