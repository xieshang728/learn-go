package tree

import "fmt"

type TreeNode struct{
	Left,Right *TreeNode
	Value int
}

func (node *TreeNode) Print(){
	fmt.Println(node.Value)
}

func createTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (node *TreeNode) SetValue(value int){
	node.Value = value
}

func PrintTreeNode(node *TreeNode){
	fmt.Println("------Tree-Node------")
	fmt.Print(node.Value)
	fmt.Println("------Tree-Node-------")
}