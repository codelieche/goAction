package main

import (
	"context"
	"fmt"
	"time"

	pb "goAction/tutorial/grpc/helloworld/protos"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {

	log.Info("Client程序开始运行")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did Not Connect RPC Server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreaterClient(conn)

	for i := 0; i < 100; i++ {

		name := fmt.Sprintf("Name %d", i)
		reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatal(err)
		} else {
			//log.Info(reply.Message)
			fmt.Println(reply.Message)
		}
		time.Sleep(time.Second)
	}

	log.Info("客户端程序执行完毕！")
}
