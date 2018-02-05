/*
切片
slice本身是没有数据的，是对底层array的一个view
slice可以向后扩展，但是不可以向前扩展
s[i]其中i不可以超越len(s), 向后扩展不可以超越底层数组cap(s)


s2 := append(s1, 22)
添加元素时，如果超越cap，系统会重新分配一个更大的底层数组
由于是指传递的缘故，必须接收appen的返回值

 */
package main

import (
	"fmt"
)

func main() {
	//sliceDemo01()

	//sliceDemo02()

	//sliceDemo03()

	sliceDemo04()

	sliceCreate()


}

func sliceDemo01() {
	fmt.Println("=== slice demo 01 ===")
	arr01 := []int{1, 3, 5, 7, 9, 11}

	fmt.Println(arr01)
	fmt.Println("arr01[:4] = ", arr01[:4])
	fmt.Println("arr01[2:4] = ", arr01[2:4])
	fmt.Println("arr01[2:] = ", arr01[2:])
	fmt.Println("arr01[:] = ", arr01[:])

	fmt.Println("\n")
}

func sliceDemo02() {
	fmt.Println("=== slice demo 02 ===")
	arr01 := []int{1, 3, 5, 7, 9, 11}

	fmt.Println("arr01 = ", arr01)
	fmt.Println("arr01[:4] = ", arr01[:4])
	s1 := arr01[:2]
	fmt.Println("s1 = ", s1)
	s2 := arr01[:]
	fmt.Println("s2 = ", s2)
	updateSlice(s2)
	fmt.Println("调用updateSlice后：s2 = ", s2)
	// 注意s1也被改掉了
	fmt.Println("调用updateSlice后：s1 = ", s1)
	fmt.Println("调用updateSlice后：arr01 = ", arr01)

}

func updateSlice(s []int) {
	// []int中括号内不加长度，它就是个slice
	// 注意这种更新切片，会把原数组中的值都改掉了
	// 想要改变数组，不要用指针，用slice即可
	s[0] = 123456
}

func sliceDemo03() {
	// 重点注意了，与py有点差异
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("=== slice demo 03 ===")
	fmt.Println("=== re slice ===")
	fmt.Println("arr: ", arr)
	s1 := arr[:4]
	fmt.Println("s1 = ", s1) // s1 =  [a b c d]
	s2 := arr[3: 5]
	// 注意，这里没报错，如果是py，应该是取不到越界的元素的
	// 但是如果尾部的边界大于cap了 就会报错
	fmt.Println("s2 = ", s2) // s2 =  [d e]

	// slice其实保存了三个值: ptr, len, cap
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 会出错
	// fmt.Printf("s2[:9] = ", s2[:9])
}

func sliceDemo04() {
	// 向切片添加元素
	fmt.Println("==== 追加元素到slice ===")
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:5]
	fmt.Println("s1 = ", s1)
	s2 := append(s1, 22)
	fmt.Println("s2 = ", s2)
	fmt.Println("执行append后s1 = ", s1)

	s3 := append(s2, 23)
	fmt.Println("s3 = ", s3)
	// 添加元素时，如果超越cap，系统会重新分配一个更大的底层数组
	// 由于是指传递的缘故，必须接收appen的返回值

	fmt.Println("\n")
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func sliceCreate() {
	// slice的创建
	fmt.Println("=== sliceCreate ===")
	var s []int  // zero value for slice is nil
	fmt.Println("s = ", s)

	for i := 0; i < 20; i++ {
		printSlice(s)
		s = append(s, 2*i + 1)
	}

	fmt.Println("s = ", s)

	s1 := []int{1, 2, 22, 33, 44, 55, 66, 77, 88, 99, 111,123, 150, 169}
	fmt.Print("s1: ")
	printSlice(s1)
	// s1: len=14, cap=14

	s2 := make([]int, 16)
	fmt.Print("s2: ")
	printSlice(s2)
	// s2: len=16, cap=16

	s3 := make([]int, 16, 32)
	fmt.Print("s3: ")
	printSlice(s3)
	// s3: len=16, cap=32

	fmt.Println("=== Copying slice ===")
	copy(s2, s1)
	fmt.Print("s2 = ", s2)
	printSlice(s2)

}
