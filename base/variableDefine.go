/*
变量定义
1. hellowWorld：简单使用
2. variableNoneValue: 定义变量不赋值
3. variableInitialValue: 定义变量且赋初始值
4. variableTypeDeduction: 变量类型推断
5. main: 主函数，main函数只能用于main包中，且只能定义一个
 */

package main

import (
	"fmt"
	"reflect"
)

func helloWorld(){
	fmt.Println("Hello Go Action")
	fmt.Printf("Hello %s", "codelieche")
	fmt.Println("\n------\n")
}

func variabelNoneValue() {
	// 定义变量，不设置值
	// Go语言比较严格，定义的变量一定要用到

	// 变量名在前面，变量类型在后面的
	// 在其它语言中，很多是类型在前，变量在后的，比如Java： int age = 18;
	var age int
	var name string
	// 打印默认值
	fmt.Printf("name: %s \t age: %d", name, age)
	// name:  	 age: 0
	fmt.Println()
	fmt.Println("=========\n")
}

func varaibalInitailValue() {
	// 定义变量设置初始值

	// 连续定义2个变量
	var year1, year2 int = 2009, 1991
	var name1, name2 string = "Golang", "Python"

	fmt.Println(name1, year1)
	fmt.Println(name2, year2)
	fmt.Println(".................\n")
}

func variableTypeDeduction() {
	// 自动推断变量类型

	// 方式1：使用var
	//var a, b, c, d = 10, 3.14, true, "Codelieche"

	// 方式2：使用冒号=
	a, b, c, d := 10, 3.14, true, "Codelieche"
	// 第一次使用变量要使用:=, 第二次就不能使用了，再使用就是重复定义变量了
	// 注意在函数外面这种是不可以的，需要用var开始
	// 对a重新赋值
	a = 2018
	fmt.Println(a, b, c, d)
	// 打印变量的类型，Python中可以直接用type
	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))
	// 2018 3.14 true Codelieche
	// 2018 int
	// 3.14 float64
	// true bool
	// Codelieche string
	fmt.Println("--------------\n")
}

func main() {
	// 调用函数
	helloWorld()

	// 变量不设置初始值
	variabelNoneValue()

	// 变量设置初始值
	varaibalInitailValue()

	// 自动推断变量类型
	variableTypeDeduction()
}
