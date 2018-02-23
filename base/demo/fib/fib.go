package main

import "fmt"

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

func main() {
	f := fibonacci()

	// 每调用一次f，获取下一个斐波那契的值
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
