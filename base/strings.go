/*
rune相当于go的char

- 使用range遍历pos，rune对
- 使用utf8.RuneCountInString获得字符数量
- 使用len获得字节长度
- 使用[]byte获得字节

其它字符串操作
- Fields, Split, Join
- Contains, Index
- ToLower, ToUpper
- Trim, TrimRight, TrimLeft
 */
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//stringDemo01()

	//stringDemo02()

	stringDemo03()

}

func stringDemo01() {
	s := "学习Golang"
	fmt.Printf("%v: len is %d\n", s, len(s))
	// 学习Golang: len is 12

	fmt.Println("=== byte ===")
	for _, b := range []byte(s) {
		fmt.Printf("%b ", b)
	}
	fmt.Println()
	for _, b := range []byte(s) {
		fmt.Printf("%o ", b)
	}
	fmt.Println()
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
}

func stringDemo02() {

	s := "学习Golang"
	fmt.Println("\n=== rune ===")
	for index, ch := range []byte(s) {
		fmt.Printf("(%d, %X)", index, ch)
	}

	fmt.Println("\nRune Count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
}

func stringDemo03() {

	s := "学习Golang"
	fmt.Println("\n=== string demo03 ===")
	for index, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", index, ch)
	}
	fmt.Println()
}