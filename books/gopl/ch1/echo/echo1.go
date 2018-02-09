package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("=== for 打印 ===")
	// 打印传递来的参数
	for i, s := range os.Args {
		fmt.Println(i, " ==> ", s)
	}
}
