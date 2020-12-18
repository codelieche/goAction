// channel的关闭
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		//fmt.Printf("Worker %d received %d\n", id, <-c)

		// 如果chan关闭了，就会持续获取到其数据类型的默认值，int是0
		// 通过ok来判断是否chan关闭了
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)

	}
}

func worker02(id int, c chan int) {
	for {
		//fmt.Printf("Worker %d received %d\n", id, <-c)

		// 如果chan关闭了，就会持续获取到其数据类型的默认值，int是0

		for n := range c {
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}
}

func closeChanel() {
	c := make(chan int)
	go worker(0, c)
	go worker02(2, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Microsecond)
	// 关闭chanel
	close(c)
	// 延时20毫秒
	time.Sleep(20 * time.Microsecond)
}

func main() {
	fmt.Println("=== close channel ===")
	closeChanel()
}
