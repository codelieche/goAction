package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("index:", r.URL)
	w.Write([]byte("Hello Index!"))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("helloworld:", r.URL)
	w.Write([]byte("Hello Golang!"))
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/", index)

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("启动服务出错：", err)
		return
	} else {
		fmt.Println("Done")
	}
}
