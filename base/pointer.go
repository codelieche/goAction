/*
指针
- go的指针不能运算
- go的指针&放在前面
- 注意&和*的使用

值传递？引用传递
- Go语言只有值传递一种方式

*/
package main

import "fmt"

func main() {
	helloPointer()

	fmt.Println("---- setValue ---")
	var a int = 10
	setValue(a)
	fmt.Printf("a：%d\n", a)

	fmt.Println("---- setValue02 ---")
	setValue02(&a)
	fmt.Printf("执行了setValue02后a：%d\n", a)

	//---- setValue ---
	//	x使用前：10
	//x重新复制后：20
	//a：10
	//---- setValue02 ---
	//指针的值：842350534880,Value：10
	//执行完后指针：842350534880, Value: 190
	//执行了setValue02后a：190

	fmt.Println("---- Swap ----")
	a, b := 100, 200
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)



}

func helloPointer() {
	var x int = 1
	var pa *int = &x
	*pa = 5
	fmt.Println(x) // 5
}

func setValue(x int) {
	fmt.Printf("x使用前：%d\n", x)
	x += 10
	fmt.Printf("x重新复制后：%d\n", x)
}

func setValue02(x *int) {
	fmt.Printf("指针的值：%d,Value：%d\n", x, *x)
	*x = 190
	fmt.Printf("执行完后指针：%d, Value: %d\n", x, *x)
}

func swap(x, y *int) {
	// 交互两个变量的值
	*x, *y = *y, *x
}