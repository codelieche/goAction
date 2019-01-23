package main

/**
运行方式：
- go run ip.go 192.168.1.123
- go run io.go  go run ip.go AB:0:0:0:0:0:FF:99
*/

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "运行方式：%s ip-address", os.Args[0])
		os.Exit(1)
	} else {
		fmt.Println("输入的IP地址名称是：", os.Args[1])
		name := os.Args[1]
		addr := net.ParseIP(name)
		if addr != nil {
			fmt.Printf("The address is %s(Type %T)\n", addr, addr)
		} else {
			fmt.Println("Invalid adress")
		}
		os.Exit(0)
	}
}
