package main

import (
	"fmt"
	"time"
)

func goroutineDemo() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for n := 0; n < 10; n++ {
				fmt.Printf("Hello from goroutine %d-%d\n", i, n)
				//runtime.Gosched() // 手动交出控制权，一般不用
			}
		}(i)
	}
	// 延时一下，要不会立刻就执行完毕了
	time.Sleep(5 * time.Microsecond)
}

func main() {
	goroutineDemo()
}
