package main

import (
	"fmt"
	"goAction/base/demo/queue"
	"goAction/base/demo/tree"
)

func main() {
	q := queue.Queue{}
	q.Push(2)
	q.Push(4)
	q.Push(6)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())


	root := tree.Node{}
	root.SetValue(10)

	fmt.Println(root)
}
