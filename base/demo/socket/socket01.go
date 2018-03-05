package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func ReceiveClientMessage(client net.Conn) {
	defer client.Close()
	fmt.Printf("开始处理client：%s的消息\n", client.RemoteAddr())
	fmt.Printf("开始处理client：%s的消息\n", client.LocalAddr())
	for {
		buf := make([]byte, 1024)
		l, err := client.Read(buf)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		msg := string(buf[:l-1])
		fmt.Printf("收到消息:%s\n", msg)

		if strings.Trim(msg, "\r") == "close" || msg == "close" {
			fmt.Println("收到消息，close")
			break
		}
		client.Write([]byte("收到消息：" + msg + "\n"))
	}
}

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer server.Close()
	fmt.Println("server is ok: ", server)
	fmt.Printf("服务段的地址：%s的\n", server.Addr())

	var i int
	i = 0
	for {
		time.Sleep(time.Second)
		fmt.Printf("==== %d ===\n", i)
		client, err := server.Accept()
		if err != nil {
			fmt.Println("出现异常：", err.Error())
		} else {
			fmt.Println("收到新的连接：", client.RemoteAddr())
			// 启动个goroutine接收客户端发来的消息
			go ReceiveClientMessage(client)
		}
		i += 1
		// 为了测试只获取5个client就关闭
		if i > 5 {
			fmt.Println("已经多余5个客户端连接了，准备关闭了")
			break
		}
	}
	fmt.Println("程序运行完毕")
}
