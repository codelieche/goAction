package main

import (
	"context"
	"net"

	pb "goAction/tutorial/grpc/helloworld/protos"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

/**
要引入pd，记得先执行：protoc --go_out=plugins=grpc:. helloworld.proto
*/

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//fmt.Println(in.Name)
	return &pb.HelloReply{Message: "Hello :" + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreaterServer(s, &server{})

	s.Serve(listen)

}
