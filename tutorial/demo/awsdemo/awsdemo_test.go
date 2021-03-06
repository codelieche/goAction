package awsdemo

import (
	"testing"
	"bytes"
)

/**
上传文件到Bucket测试
*/

var bucketName = "myBucketName"
func TestUploadFileToS3(t *testing.T) {
	// 先准备好文件
	buff := bytes.NewBufferString("Hello This is Test")
	url, err := UpljoadBinFileToS3(buff, bucketName,
		"test/1234.txt", "text/html")
	if err != nil {
		t.Error("上传图片失败：", err.Error())
	} else {
		t.Log("上传文件成功:", url)
	}
}

// 删除s3对象测试
func TestDeleteS3Object(t *testing.T) {
	var bucketName = "test"
	var keyName = "test/1234.txt"
	if result, err := DeleteS3Object(bucketName, keyName); err != nil {
		t.Error("删除对象失败：", err.Error())
	} else {
		if result {
			t.Log("删除对象成功")
		} else {
			t.Error("删除对象失败！但是没报错")
		}
	}
}
