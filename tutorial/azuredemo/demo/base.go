package demo

import (
	"context"
	"fmt"
	"goAction/tutorial/azuredemo/config"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func handleErrors(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func GetContainerURL(containerName string) (containerURL azblob.ContainerURL, err error) {
	// 实例化变量
	var (
		credential *azblob.SharedKeyCredential // 操作的凭证
	)

	// 第1步：获取到配置, 校验数据
	config := config.Config
	if config.AccountName == "" {
		log.Println("config.AccountName为空")
		os.Exit(1)
	}
	if containerName == "" {
		log.Println("container name为空")
		os.Exit(1)
	}

	// 2. 实例化容器的client
	// 2-1: create a credential: 实例化操作的凭证
	if credential, err = azblob.NewSharedKeyCredential(config.AccountName, config.AccessKeySource); err != nil {
		log.Println("实例化连接的凭证出错：", err.Error())
		log.Panic(err)
		os.Exit(1)
	}

	// 2-2：实例化PipLine
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// 2-3：实例化ContainerUrl
	URL, _ := url.Parse(
		fmt.Sprintf(
			"https://%s.blob.core.windows.net/%s",
			config.AccountName,
			containerName))

	containerURL = azblob.NewContainerURL(*URL, p)
	return containerURL, err
}

// 上传文件，列出容器中的文件
func AzureBlobBase(containerName string) (err error) {

	containerUrl, err := GetContainerURL(containerName)
	if err != nil {
		log.Println("获取containerURL出错")
		return
	}

	log.Println("containerUrl:", containerUrl)

	// 上传文件
	blobUrl := containerUrl.NewBlockBlobURL("golang.html")
	rsp, err := azblob.UploadBufferToBlockBlob(context.Background(), []byte("This is Golang Test!\n"), blobUrl, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16,
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "text/html",
		},
	})
	if err != nil {
		log.Println("上传图片出错")
		log.Print(err)
		os.Exit(1)
	} else {
		log.Println("上传文件成功")
		log.Println(rsp.Date(), rsp.LastModified(), rsp.ETag())
	}

	// 列出容器中的Blob
	fmt.Printf("\n\n")

	log.Println("开始列出blob")
	maxCount := 200
	for marker := (azblob.Marker{}); marker.NotDone(); {
		listBlob, err := containerUrl.ListBlobsFlatSegment(context.Background(), marker, azblob.ListBlobsSegmentOptions{})

		handleErrors(err)

		marker = listBlob.NextMarker

		for _, blobInfo := range listBlob.Segment.BlobItems {
			maxCount -= 1
			if maxCount < 0 {
				goto END
			}
			fmt.Println("Blob name:", blobInfo.Name, *blobInfo.Properties.ContentType)
		}
	}
END:
	return nil
}
