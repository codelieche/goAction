package main

import (
	"fmt"
	"goAction/base/demo/retriever/mock"
	"goAction/base/demo/retriever/real"
)

type Retriever interface {
	// 在interface里面，不需要加func
	// 里面本身就是函数
	Get(url string) string
}

func download(r Retriever) string {
	// 使用者：download
	return r.Get("http://www.codelieche.com")
}

func main() {
	var r Retriever
	fmt.Println("=== Mock ====")
	r = mock.Retriever{"This is a facke codelieche.com"}
	fmt.Println(download(r))
	fmt.Println("==== Real ====")
	r2 := real.Retriever{}
	fmt.Println(len(download(r2)))

}
