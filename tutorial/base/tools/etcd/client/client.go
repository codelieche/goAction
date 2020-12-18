package main

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端的配置
	config = clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 20 * time.Second, // 超时时间
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println("连接错误")
		fmt.Println(err)
		return
	} else {
		defer client.Close()
		fmt.Println(client)
	}
}
