package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
	}

}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	// 创建10个worker数组
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// 给chan中加入值
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		// 收
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		// 收
		<-workers[i].done
	}

	//time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}
