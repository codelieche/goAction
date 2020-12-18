/*
内建变量类型
- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, unitptr
  1. 加了u就是无符号类型的
  2. 不加长度，就根据操作系统来的，32位的操作系统就是32位，64位的操作系统就是64位的

- byte, rune(rune是go语言的char类型)
- float32, float64, complex64, complex128

1. defineConsts: 定义常量

 */

package main

import (
	"math"
	"fmt"
)

func defineConsts() {
	// 定义常量
	const name = "codelieche"
	const x, y = 3, 4
	z := int(math.Sqrt(x*x + y*y))
	fmt.Println(name, x, y, z)
}

func defineEnums() {
	// 定义枚举
	const (
		c = 0
		java = 1
		python = 2
		golang = 3
		js = 4
	)
	fmt.Println(c, java, python, golang, js)

}

func useIota() {
	fmt.Println("==== iota ====")
	const (
		a1 = iota
		b1
		c1
		d1
	)
	fmt.Println(a1, b1, c1, d1)

	fmt.Println("==== iota ====")

	// iota是个自增值的种子
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	// 定义常量
	defineConsts()

	// 枚举
	defineEnums()

	// iota
	useIota()
}