/*
函数

函数式编程：函数是一等公民！

- 返回值类型写在最后面
- 可返回多个值
- 函数可以做为参数的
- 注意：golang函数是没有默认参数，可选参数的！！！！

 */

package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	// 无返回值
	doSomething(10, 20)

	// 有返回值
	fmt.Println("--------")
	fmt.Println(doSomething02(10, 20))

	// apply
	fmt.Println(apply(doSomething02, 10, 20))
	// 调用函数:main.doSomething02，传递的参数是(10, 20)

	fmt.Println("------- sum -------")
	result := sumArgs(1, 2, 3, 4, 5)
	fmt.Printf("1, 2, 3, 4, 5 = %d", result)


}

func doSomething(a, b int) {
	fmt.Println(a, b)
	fmt.Println("Do something")
}

func doSomething02(a, b int) int {
	fmt.Println(a, b)
	fmt.Println("Do something 02")
	return a + b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("调用函数:%s，传递的参数是(%d, %d)\n", name, a, b)
	return op(a, b)
}

func sumArgs(values ...int) int {
	// 可变参数
	sum := 0

	// 方式1：只要索引，取值通过索引获取
	//for i := range values {
	//	sum += values[i]
	//}

	// 方式2：遍历索引和值
	for index, value := range values {
		fmt.Println(index, value)
		sum += value
	}

	// 方式2： 忽略索引，只要值
	for _, value := range values {
		fmt.Println(value)
	}

	return sum
}


