package main

import (
	"fmt"
	"goAction/base/demo/queue"
	"goAction/base/demo/tree"
)

func main() {
	fmt.Print("=== Queue ===")
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

	fmt.Println("\n=== Queue2 ===")
	q2 := queue.Queue2{}
	q2.Push("Good")
	q2.Push("Hello")
	q2.Push(2)
	fmt.Println(q2.Pop())
	fmt.Println(q2.Pop())
	fmt.Println(q2.IsEmpty())
	fmt.Println(q2.Pop())
	fmt.Println(q2.IsEmpty())
}
