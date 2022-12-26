package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{5, nil}}}}}
	ret := deleteDuplicates(head)
	show(ret)

	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{5, nil}}}}}
	ret = deleteDuplicates_83(head)
	show(ret)
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			x := p.Next.Val
			for p.Next != nil && p.Next.Val == x {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
func deleteDuplicates_83(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
func show(node *ListNode) {
	if node == nil {
		fmt.Println("---------------------------")
		return
	}
	fmt.Println(node.Val)
	show(node.Next)
}
