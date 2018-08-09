// dup1 输出标准输入中出现次数大于1的行，前面的是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map类似于python中的dict {"string": int}
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if len(text) > 0 {
			//fmt.Println(text)
			counts[text]++
		}else {
			break
		}
	}

	fmt.Println(" ==== 统计结果 ===")

	// 注意：忽略 input.Err() 中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
