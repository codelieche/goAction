package main

import (
	"log"

	"codelieche.com/server"
)

func main() {
	log.Println("monitor程序开始运行")
	server.Run()
	log.Println("monitor程序执行完毕")
}
