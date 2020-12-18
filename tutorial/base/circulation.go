/*
循环
 */
package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"time"
)

func main() {
	fmt.Print("1 加到 100 = ")
	fmt.Println(sum1To100()) // 5050


	fmt.Println(
		"二进制转换",
		convertToBin(10),
		convertToBin(20),
		convertToBin(50),
	)

	fmt.Println("----- 打印文件 ----")
	printFile("./README.md")

	fmt.Println("---- for ever ----")
	forEver()
	fmt.Println("---- for ever done ----")
}


func sum1To100() int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	// 打印文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 关闭文件
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func forEver() {
	// 来个死循环
	count := 0
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)

		// 只输出100次
		count++
		if count > 100 {
			break
		}
	}
}
