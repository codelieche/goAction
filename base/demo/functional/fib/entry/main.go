package main

import (
	"bufio"
	"fmt"
	"goAction/base/demo/functional/fib"
	"io"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()

	// 每调用一次f，获取下一个斐波那契的值
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("\n=== printFileContents ===")
	var g fib.IntGen = fib.Fibonacci()
	g.Test123(6)

	printFileContents(g)

}
