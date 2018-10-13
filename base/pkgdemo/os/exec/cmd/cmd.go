package cmd

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd  *exec.Cmd
	)
	// 实例化Cmd
	cmd = exec.Command("/bin/bash", "-c", "echo `date`;ls;pwd")

	// 执行命令: 只捕获错误
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("执行完毕")
	}

	// 执行命令：捕获子进程的输出(pipe)
	cmd = exec.Command("/bin/bash", "-c", "echo `date +\"%F %T\"`;ls;pwd")
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}else{
		// 打印进程的输出
		fmt.Println(output)
		fmt.Println(string(output))
	}

}
