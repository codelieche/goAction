package main

import (
	"fmt"
	"goAction/base/demo/retriever/mock"
	"goAction/base/demo/retriever/real"
	"time"
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

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case real.Retriever:
		fmt.Println("Real: UserAgent", len(v.UserAgent))
	case *real.Retriever:
		fmt.Println("*Real: UserAgent", len(v.UserAgent))
	default:
		fmt.Println("default")
	}

}

func typeAssert(r Retriever) {
	// type assertion 类型断言
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
}

func main() {
	var r Retriever
	fmt.Println("=== Mock ====")
	r = mock.Retriever{"This is a facke codelieche.com"}
	fmt.Println(download(r))
	fmt.Printf("%T %v\n", r, r)

	fmt.Println("\n==== Real ====")
	r2 := real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Println(len(download(r2)))
	fmt.Printf("%T %v\n", r2, r2)
	fmt.Printf("%T %v\n", &r2, &r2) // *real.Retriever &{Mozilla/5.0 1m0s}

	fmt.Println("\n=== inspect ===")
	inspect(r)
	inspect(&r2)
	inspect(r2)

	fmt.Println("\n=== typ assert ===")
	fmt.Print("r:\t")
	typeAssert(r)
	fmt.Print("r2:\t")
	typeAssert(r2)

}
