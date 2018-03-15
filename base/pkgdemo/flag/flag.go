package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "默认名", "请输入用户名")
	password := flag.String("password", "pwd", "请输入密码")
	flag.Parse()

	// go run ./flag.go --username=用户 --password=abc1234567
	fmt.Println(username, password)
	fmt.Println("值：", *username, *password)
}
