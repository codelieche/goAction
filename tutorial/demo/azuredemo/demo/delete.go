package demo

import (
	"context"
	"log"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func DeleteBlobFromAzure(containerName string, blobName string) (err error) {
	var (
		containerURL azblob.ContainerURL
		blobURL      azblob.BlockBlobURL
	)
	// 第1步：获取到containerURL
	if containerURL, err = GetContainerURL(containerName); err != nil {
		return
	}

	// 第2步：实例化blobURL
	blobURL = containerURL.NewBlockBlobURL(blobName)

	// 第3步：操作删除
	ctx := context.Background()
	if resp, err := blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{}); err != nil {
		log.Print("删除失败")
		return err
	} else {
		log.Println("删除成功：", resp.RequestID(), resp.ErrorCode())
	}
	return err
}
