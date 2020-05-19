package demo

import (
	"context"
	"fmt"
	"goAction/tutorial/azuredemo/config"
	"io"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func DownloadBlobFromAzure(containerName string, blobName string) (err error) {
	// 实例化变量
	var (
		credential   *azblob.SharedKeyCredential // 操作的凭证
		containerURL azblob.ContainerURL         // 容器的连接
		blobURL      azblob.BlockBlobURL         // Blob的连接
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

	// 第3步：下载文件
	// 3-1： 实例化BlobURL
	blobURL = containerURL.NewBlockBlobURL(blobName)

	// 3-2: 开始下载
	ctx := context.Background()
	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)

	if err != nil {
		return err
	}

	// 3-3: 读取下载的响应
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	//downloadedData := bytes.Buffer{}

	//_, err = downloadedData.ReadFrom(bodyStream)
	fileName := fmt.Sprintf("/tmp/azure_download_%s", blobName)
	if file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 777); err != nil {
		log.Printf("创建文件%s出错:%s\n", blobName, err.Error())
	} else {
		if _, err := io.Copy(file, bodyStream); err != nil {
			log.Println("io.Copy数据出错！")
		}
	}
	return err
}
