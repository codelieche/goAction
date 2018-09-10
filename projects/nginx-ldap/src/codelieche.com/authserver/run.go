package authserver

import (
	"log"
	"net/http"
)

func Run() {
	log.Print("Web Server 开始运行")

	// 1. 处理auth校验：GET/POST
	// 1-1：auth的处理器
	authHandler := http.HandlerFunc(HandleAuth)

	// 1-2: 给auth添加session相关的中间件
	http.Handle("/auth", checkSessionMiddleWare(authHandler))

	// 2. 登录用户的相关操作

	// 3. 启动web服务
	addr := ":9000"
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(addr, err.Error())
	} else {
		msg := "web服务执行完毕！"
		log.Println(msg)
	}

}
