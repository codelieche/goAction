package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8...
// a, b
//    a, b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func (g intGen) Test123(n int) {
	fmt.Println("=== Test123 ===")
	fmt.Println(g, "\t", n)
	fmt.Println("=== Test123 End ===")
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	// 每调用一次f，获取下一个斐波那契的值
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("\n=== printFileContents ===")
	var g intGen = fibonacci()
	g.Test123(6)

	printFileContents(g)

}
