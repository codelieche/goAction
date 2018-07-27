package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
	"goAction/tutorial/rpcdemo"
)

func main(){
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":4567")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		log.Println(conn, conn.LocalAddr(), conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
