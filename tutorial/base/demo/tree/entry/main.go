package main

import (
	"fmt"
	"goAction/base/demo/tree"
);

func main()  {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	fmt.Println(root)
	fmt.Println("====== Root ====")
	root.Right.Left.SetValue(4)
	root.Print()

	fmt.Println("\n==== 遍历Value ===")
	root.Traverse()
}
