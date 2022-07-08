package main

import (
	"container/heap"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeBool []*ListNode

func (n NodeBool) Len() int {
	return len(n)
}
func (n NodeBool) Less(i, j int) bool {
	return (*n[i]).Val < (*n[j]).Val
}
func (n NodeBool) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NodeBool) Push(x any) {
	en := x.(*ListNode)
	*n = append(*n, en)
}

func (n *NodeBool) Pop() any {
	l := len(*n)
	old := (*n)
	node := old[l-1]
	*n = old[0 : l-1]
	return node
}

// 注意 Push 前判断是否为空
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	nodeBool := &NodeBool{}
	for idx, _ := range lists {
		if lists[idx] != nil {
			heap.Push(nodeBool, lists[idx])
		}
	}
	p1 := head
	for nodeBool.Len() > 0 {
		minNode := heap.Pop(nodeBool).(*ListNode)
		if minNode.Next != nil {
			heap.Push(nodeBool, minNode.Next)
		}
		p1.Next = minNode
		p1 = p1.Next
	}
	return head.Next
}

func main() {
	nodeBool := &NodeBool{}
	nodeBool.Push(&ListNode{1, nil})
	nodeBool.Push(&ListNode{2, nil})
	nodeBool.Push(&ListNode{3, nil})
	nodeBool.Push(&ListNode{4, nil})
	fmt.Println(nodeBool.Len())
	fmt.Println((nodeBool.Len() - 1 - 1) / 2)
	fmt.Println(nodeBool)
}
