package demo

import "testing"

func TestDownloadBlobFromAzure(t *testing.T) {
	containerName := "test"
	blobName := "golang.html"

	if err := DownloadBlobFromAzure(containerName, blobName); err != nil {
		t.Error(err)
	} else {
		t.Log("下载文件成功，请去tmp查询！")
	}
}
