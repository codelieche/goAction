package main

/**
go run mask.go 192.168.6.106 22 32
IP地址是： 192.168.6.106
地址长度是： 32
网络位长度是： 22
子网掩码(hex)是： fffffc00
网络地址是： 192.168.4.0

*/
import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "本脚本使用方式是：%s ip-addr 网络位长度 字节长度\n", os.Args[0])
		os.Exit(1)
	} else {
		// 开始处理ip
		dotAddr := os.Args[1]
		ones, _ := strconv.Atoi(os.Args[2])
		bits, _ := strconv.Atoi(os.Args[3])

		addr := net.ParseIP(dotAddr)

		if addr == nil {
			fmt.Printf("地址错误：", dotAddr)
			os.Exit(1)
		}

		// 获取网关等
		mask := net.CIDRMask(ones, bits)
		network := addr.Mask(mask)
		fmt.Println("IP地址是：", addr.String())
		fmt.Println("IP地址长度是：", bits)
		fmt.Println("网络位长度是：", ones)
		fmt.Println("子网掩码(hex)是：", mask.String())
		fmt.Println("网络地址是：", network.String())
	}
}
