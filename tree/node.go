package main

import (
	"fmt"
	"learn-go/tree/entity"
)

type MyNode struct {
	node *tree.TreeNode
}

func (node *MyNode)printMyNode(){
	fmt.Println("====组合=====")
	fmt.Println(node)
}

func main(){
	var root tree.TreeNode
	root = tree.TreeNode{Value: 5}
	root.Left = &tree.TreeNode{Value:10}
	root.Right = new(tree.TreeNode)
	root.Print()
	root.SetValue(111)
	root.Print()
	tree.PrintTreeNode(&root)
	tree.TraverseNode(&root)


	node := MyNode{&root}
	node.printMyNode()


}