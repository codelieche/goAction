package main

import (
	"fmt"
	"time"
)

func chanDemo01() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	// 把1，2送给channel
	c <- 1
	c <- 2

	time.Sleep(time.Millisecond)
}

func chanDemo02() {
	c := make(chan string)
	go worker(c)

	c <- "字符1"
	c <- "abcd"

	time.Sleep(time.Millisecond)
}

func worker(c chan string) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func chanDemo03() {
	c := make(chan int)
	// 启动10个worker
	for i := 0; i < 10; i++ {
		go worker02(i, c)
	}

	// 传递100个数字到c中
	for n := 0; n < 100; n++ {
		c <- n
		time.Sleep(time.Microsecond)
	}

	time.Sleep(time.Millisecond)
}

func worker02(id int, c <-chan int) {
	// 关于参数：
	// <-chan：这样表示c只能发数据，不可以往里面加数据
	// chan <- int: 表示只可加入数据到chan中，不可以取
	for {
		n := <-c
		fmt.Printf("Workder %d received %d \n", id, n)
	}
}

func main() {
	fmt.Println("=== channel Demo 01 ===")
	chanDemo01()

	fmt.Println("\n=== channel Demo 02 ===")
	chanDemo02()

	fmt.Println("\n=== channel Demo 03 ===")
	chanDemo03()

}
