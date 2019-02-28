package smallweb

import (
	"fmt"
	"log"
	"net/http"
)

/*
启动个Web服务：
- `/`: 首页
- `/api`: 接收GET、POST方法传参
*/

func RunServer() {
	log.Println("启动Web Server")
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	http.HandleFunc("/", webRoute)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(addr, err.Error())
	} else {
		msg := "Web Server执行完毕"
		log.Println(msg)
	}
}
