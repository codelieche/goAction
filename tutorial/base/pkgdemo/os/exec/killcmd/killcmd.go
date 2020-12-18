package main

import (
	"context"
	"os/exec"
	"time"
	"fmt"
)

/**
启动个命令然后3后，杀掉这个进程
 */

type ExecuteResult struct {
	// 执行结果
	Id int          // 执行的编号
	Start time.Time // 开始执行的时间
	End time.Time 	// 结束时间
	Cmd string      // 执行的命令
	Success bool    // 执行是否成功
	Result string	// 执行命令的输出结果
}


func main(){
	var (
		// 执行命令的上下文
		ctx context.Context
		//	执行命令上下文的取消函数
		// 上下文和取消函数是成对的
		cancelFun context.CancelFunc

		cmd *exec.Cmd

		// 执行命令：捕获的异常、输出的结果
		err error
		output []byte

		// 执行结果的chan
		resultChan chan *ExecuteResult
	)

	// 得到上下文和取消函数
	ctx, cancelFun = context.WithCancel(context.TODO())

	// 创建结果队列
	resultChan = make(chan *ExecuteResult, 100)

	go func() {
		cmdstr := "echo `date`;sleep 5;echo Hello!"
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", cmdstr)

		// 执行结果
		result := ExecuteResult{
			Id: 1,
			Start: time.Now(),
			Cmd: cmdstr,
		}

		// 当执行命令的上下文关闭的时候，会杀掉进程
		// ctx.Done():  kill pid, 杀掉子进程
		output, err = cmd.CombinedOutput()
		result.End = time.Now()
		if err != nil {
			fmt.Println("执行命令出错：", err)
			result.Success = false
			result.Result = err.Error()
		}else{
			// fmt.Println("执行命令成功：", cmdstr)
			// fmt.Println(string(output))

			result.Success = true
			result.Result = string(output)
		}

		// 发送结果到队列中
		resultChan <- &result
	}()

	// 延时3秒
	time.Sleep(3 * time.Second)

	// 取消上下文
	cancelFun()

	// 获取结果队列数据
	//result := &ExecuteResult{}
	result := <- resultChan
	fmt.Println(result)

	fmt.Println("程序执行完毕！")
}
