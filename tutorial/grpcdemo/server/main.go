package main

import (
	"context"
	"goAction/tutorial/grpcdemo/proto"

	"net"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

type CheckPermissionService struct{}

func (s *CheckPermissionService) CheckPermission(ctx context.Context, r *proto.CheckRequest) (*proto.CheckResponse, error) {
	return &proto.CheckResponse{Status: true, Message: r.GetUsername() + r.GetPermission()}, nil
}

const ADDR = ":4567"

func main() {
	log.Info("开始执行Server端")
	server := grpc.NewServer()
	proto.RegisterPermissionCheckServer(server, &CheckPermissionService{})

	lintner, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		server.Serve(lintner)
	}
}
