/*
条件语句

 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("--- if ---")
	ifDemo01(28)
	ifDemo01(16)
	ifDemo01(18)
	fmt.Println("------------\n")

	ifDemo02("README.md")
	ifDemo02("README02.md")

	fmt.Println("\n--- switch ---")
	fmt.Println("11 + 20 = ", switchDemo01(11, 20 , "+"))
	fmt.Println("11 * 20 = ", switchDemo01(11, 20 , "*"))
	//fmt.Println("11 ** 20 = ", switchDemo01(11, 20 , "**"))

	fmt.Println("99分：", switchDemo02(99))
	fmt.Println("67分：", switchDemo02(67))
	fmt.Println("85分：", switchDemo02(85))
}

func ifDemo01(age int) {
	// if 条件是不需要用括号的
	if age < 18 {
		fmt.Printf("年龄(%d)小于18\n", age)
	}else if age == 18{
		fmt.Printf("年龄(%d)等于18\n", age)
	}else{
		fmt.Printf("年龄(%d)大于18\n", age)
	}
}

func ifDemo02(filename string) {
	// 读取文件

	contents, err := ioutil.ReadFile(filename)
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	if err != nil {
		fmt.Println("读取文件出错：")
		fmt.Println(err.Error())
		return
	} else {
		// 读取文件成功
		fmt.Printf("\n读取文件成功：\n%s\n", contents)
	}

}

func switchDemo01(a, b int, flag string) int {
	// switch会自动break，除非使用fallthrough
	var result int
	switch flag {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("您传递的符号我不能处理:" + flag)
	}

	// 返回结果
	return result
}

func switchDemo02(score int) string {
	grade := ""
	switch {
	case  score < 0 || score > 100:
		panic(fmt.Sprintf("错误的分数：%d", score))
	case score < 60:
		grade = "E"
	case score < 70:
		grade = "D"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	case score <= 100:
		grade = "A"
	}
	return grade
}


