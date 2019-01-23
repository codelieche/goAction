package main

import (
	"context"
	pb "goAction/tutorial/grpc/helloworld/protos"
	"net"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

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
