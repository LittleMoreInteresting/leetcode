package tree

import "fmt"

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func MockTree() *TreeNode {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}
	return t
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	fmt.Println(tree.Data)
	MidOrder(tree.Right)
}

func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Println(tree.Data)
}
