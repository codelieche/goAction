## goroutine

### 协程 Coroutine
> 一个进程中有多个线程，一个线程中可以有多个协程。

- 轻量级"线程"
- 非抢占式多任务处理，由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个线程可以在一个或多个线程上运行

**goroutine可能的切换点**
- I/O,select
- channel
- 等待锁
- 函数调用(有时)
- runtime.Gosched()

**注意：**只是参考，不能保证切换，不能保证在其他地方不切换。

**执行体:**  
> 执行体是个抽象的概念，在操作系统层面有个概念与之对应，比如操作系统自己长官的进程(process)、进程内的线程(thread)
以及进程内的协程(coroutine,也叫轻量级的线程)。

**执行体之间通信：**在Go语言中内置了消息队列的支持，只不过它叫通道(channel)。两个`goroutine`之间可以通过通道类进行交互。

