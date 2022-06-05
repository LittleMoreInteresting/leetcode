package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
	reorderList(head)
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

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

func reorderList(head *ListNode) {
	st := []*ListNode{head}
	for node := head.Next; node != nil; node = node.Next {
		st = append(st, node)
	}
	
	node := head
	idx := len(st) - 1
	for node != nil {
		lastNode := st[idx]
		idx--
		next := node.Next
		if lastNode == next || lastNode.Next == next {
			lastNode.Next = nil
			break
		}
		node.Next = lastNode
		lastNode.Next = next
		node = next
	}
}