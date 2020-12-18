package main

import (
	"context"

	"github.com/codelieche/goAction/tutorial/demo/grpcdemo/proto"

	"fmt"

	"time"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

const ADDR = ":4567"

func main() {
	fmt.Println("开始执行客户端")
	conn, err := grpc.Dial(ADDR, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
		return
	} else {
		defer conn.Close()

		client := proto.NewPermissionCheckClient(conn)

		response, err := client.CheckPermission(context.Background(),
			&proto.CheckRequest{Username: "codelieche", Permission: "can_admin_website"})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Response: %s \n", response)
		}

		// 延时10秒后再次执行查看权限
		time.Sleep(10 * time.Second)
		if response, err := client.CheckPermission(
			context.Background(),
			&proto.CheckRequest{
				Username:   "codelieche",
				Permission: "can_admin_website",
			}); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Response: %s", response)
		}
	}
}
