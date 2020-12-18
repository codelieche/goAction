package main

import (
	"fmt"
	"path"
)

func main() {
	p := path.Base(".")
	fmt.Println(p)
	fmt.Println(path.Dir("./"))

	fmt.Println("\n=== 判断path是否是绝对路径 ===")

	fmt.Println(path.IsAbs("./tmp"))
	fmt.Println(path.IsAbs("/tmp"))

	fmt.Println("\n=== 组合path ===")
	fmt.Println(path.Join("./tmp", "file01.txt"))

	fmt.Println("\n=== 分割path ===")
	fmt.Println(
		path.Split("./tmp/log.txt"),
	)

	fmt.Println("\n=== 返回path目录 ===")
	fmt.Println(
		path.Dir("./tmp/log23456.txt"),
	)

	fmt.Println("\n=== 返回path拓展名 ===")
	fmt.Println(
		path.Ext("./tmp/log23456.txt"),
	)

	fmt.Println("\n=== path Match ===")
	fmt.Println(
		path.Match("tmp/log*.txt", "tmp/log23456.txt"),
	)
	fmt.Println(
		path.Match("tmp/log", "tmp/log23456.txt"),
	)

}
