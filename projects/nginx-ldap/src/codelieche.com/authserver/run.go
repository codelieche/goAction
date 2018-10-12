package authserver

import (
	"log"
	"net/http"
)

func Run() {
	log.Print("Web Server 开始运行")

	// 1. web router
	// 1-1：webRoute的处理器
	webRouterHandler := http.HandlerFunc(webRooter)

	// 1-2: 给webRouterHandler添加session相关的中间件
	http.Handle("/", checkSessionMiddleWare(webRouterHandler))

	// 如果不用中间件，那么直接使用函数
	//http.HandleFunc("/", webRooter)

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
