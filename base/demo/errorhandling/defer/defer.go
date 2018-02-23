package main

import (
	"bufio"
	"fmt"
	"goAction/base/demo/functional/fib"
	"os"
)

func tryDefer() {
	// 先进后出：实际输出顺序是：3,4, 2, 1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
}

func tryDefer02() {
	// 先进后出：实际输出顺序是：3, 2, 1
	// 【4不会输出，虽然抛出了错误，但是2，1还是输出的o】
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("Error occurred")
	fmt.Println(4)

}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 建了文件，一定要关掉
	defer file.Close()

	writer := bufio.NewWriter(file)

	// 如果不刷新，缓存区的内容不一定到文件中
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	fmt.Println("=== tryDefer ===")
	tryDefer()
	fmt.Println("\n=== tryDefer02 ===")
	//tryDefer02()

	writeFile("tmp/fib.txt")

}
