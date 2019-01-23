package objects

import (
	"goAction/books/src/objects-storage/settings"
	"log"
	"net/http"
)

/**
运行对象存储的Server端
*/

func Run() {
	log.Println("程序开始运行")

	http.HandleFunc("/", handlerRouter)

	if err := http.ListenAndServe(settings.ADDRESS, nil); err != nil {
		log.Println("启动对象存储服务失败：", err)
		return
	}
	log.Println("程序执行完毕！")
}
