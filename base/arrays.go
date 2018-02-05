/*
数组

[5]int和[10]int是不同类型
调用func f(arr[10]int)会拷贝数组: 大部分语言可不是这样的

在python中，如果函数传递个数组，在函数内改变了数组：，数组是会改变的，但是go不会
>>> x = [1, 2, 3]
>>> def f(i):
...     i[0] = 100
...     print(i)
...
>>> x
[1, 2, 3]
>>> f(x)
[100, 2, 3]
>>> x
[100, 2, 3]
 */

package main

import "fmt"

func main() {
	arrayDemo01()

	arrayDemo02()

	//arrayPrint01()
}


func arrayDemo01() {
	var arr1 [5]int
	arr2 := [3]int{2, 4, 6}
	arr3 := [...]int{1, 3, 5, 7, 8, 10, 11}

	arr1[0] = 100
	arr2[0] = 100
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// 注意go中数组是值类型，在fun中改变数组
	fmt.Println("arr2:", arr2)
	changeArr(arr2) // 注意changeArr参数设置的是[3]int
	fmt.Println("arr2:", arr2)
	// 通过输出可以看出，changeArr没有改变arr2的值
	fmt.Println("\n")

	fmt.Println("=== 如果想改变数组的值呢? ====")
	changeArr02(&arr2) // 注意changeArr参数设置的是*[3]int
	// 传递指针，会修改o
	fmt.Println("arr2:", arr2)
	// 通过输出可以看出，changeArr没有改变arr2的值
	fmt.Println("\n")

}

func arrayDemo02() {
	var grid [2][2]int
	fmt.Println(grid)
	// [[0 0] [0 0]]
}

func arrayPrint01() {
	// 数组遍历
	arr01 := []string{"a", "b", "c", "d"}
	fmt.Println("===== index:value =====")
	for index, value := range arr01 {
		fmt.Println(index, "===>", value)
	}

	fmt.Println("===== index =====")
	// 有两种方式，一：只传一个值接收range的值，二：用下划线忽略value值
	//for index := range arr01 {
	for index, _ := range arr01 {
		fmt.Println("index: ", index)
	}

	// 忽略索引，只要值
	fmt.Println("===== value =====")
	for _, value := range arr01 {
		fmt.Println("Value: ", value)
	}

}


func changeArr(a [3]int) {
	// 传递数组过来，是会复制下，函数内部修改不会改变数组的值
	a[0] = 123456
	fmt.Println("把数组a的第一个元素改成：123456")
	fmt.Println(a)
	fmt.Println("changeArr函数执行完毕")
}

func changeArr02(a *[3]int) {
	// 传递数组的指针过来，一般不用这种方式的
	a[0] = 123456
	fmt.Println("把数组a的第一个元素改成：123456")
	fmt.Println(a)
	fmt.Println("changeArr函数执行完毕")
}

