package main

import (
	"fmt"
)

type treeNode struct {
	// 定义个树，树有左右，左右各自是个treeNode
	// 注意：加入treeNode放在其它包中，我们想在main中使用，首字母必须大写o
	value int
	// 指针o
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node treeNode) run() {
	// 不加指针，是值传递，里面修改了值，外面也收不到
	fmt.Println("执行run方法：node.value", node.value)
}

func (node *treeNode)addValue(n int){
	// 使用指针作为方法的接收者
	// 只有使用指针才可以改变结构内容
	// 传递指针，里面修改了元素的值，外面能接收到
	// nil指针也可以调用方法！
	fmt.Println("node value is: ", node.value)
	node.value += n
	fmt.Println("执行allValue + ", n)
}

func (node *treeNode) setValue(value int){
	// 给node赋值
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.value = value
}

func (node *treeNode) traverse(){
	// 遍历node
	if node == nil {
		return
	}
	// go语言这里不用判断left或者right是不是nil
	node.left.traverse()
	node.print()
	node.right.traverse()

}


func main() {
	var root  treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right =&treeNode{5, nil,nil}
	root.right.left = new(treeNode)
	root.right.left.value = 4
	root.left.right = createNode(2)

	fmt.Println(root)

	nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root}, // 注意这里也要逗号
	}
	fmt.Println(nodes)

	// 执行run方法
	fmt.Println("=== run ===")
	root.run()
	fmt.Println("=== addValue + 5 ===")
	fmt.Println("执行前：root.value = ", root.value)
	root.addValue(5)
	fmt.Println("执行后：root.value = ", root.value)

	fmt.Println("=== setValue(3) ===")
	root.setValue(3)
	fmt.Println("执行后：root.value = ", root.value)

	fmt.Println("\n=== traverse ===")
	root.traverse()
}