package config

import (
	"log"
	"os"
)

/**
获取配置
*/
var Config = &AzureConfig{
	ContainerSource: "source",
	ContainerTarget: "target",
}

type AzureConfig struct {
	// 自行规定环境变量名称
	AccountName      string // 环境变量: AZURE_STORAGE_ACCOUNT_NAME
	AccessKey        string // 环境变量：AZURE_STORAGE_ACCESS_KEY
	AccessKeySource  string // 环境变量：AZURE_STORAGE_SOURCE_ACCESS_KEY
	AccessKeyTareget string // 环境变量：AZURE_STORAGE_TARGET_ACCESS_KEY
	ContainerSource  string // 原容器名字
	ContainerTarget  string // 目标容器的名字
}

func (c *AzureConfig) ParseValue() {
	if value := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"); value != "" {
		c.AccountName = value
	} else {
		log.Println("请设置Config的AccountName")
	}

	if connectionKey := os.Getenv("AZURE_STORAGE_ACCESS_KEY"); connectionKey == "" {
		log.Println("请配置：AZURE_STORAGE_ACCESS_KEY环境变量")
		os.Exit(1)
	} else {
		c.AccessKey = connectionKey
		c.AccessKeySource = connectionKey
		c.AccessKeyTareget = connectionKey
	}

	// 接下来检查默认的源容器和目标容器的key，如果不存在就使用connectionKey
	if value := os.Getenv("AZURE_STORAGE_SOURCE_ACCESS_KEY"); value != "" {
		c.AccessKeySource = value
	}
	if value := os.Getenv("AZURE_STORAGE_TARGET_ACCESS_KEY"); value != "" {
		c.AccessKeyTareget = value
	}
}

func init() {
	log.Println("azuredemo/config init running")
	Config.ParseValue()
}
