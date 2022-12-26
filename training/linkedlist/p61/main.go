package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ret := rotateRight(head, 2)
	show(ret)
	head = &ListNode{1, nil}
	ret = rotateRight(head, 1)
	show(ret)
}
func show(node *ListNode) {
	if node == nil {
		fmt.Println("---------------------------")
		return
	}
	fmt.Println(node.Val)
	show(node.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := ListLen(head)
	if l < 2 {
		return head
	}
	for i := 0; i < k%l; i++ {
		pre, end := LastNode(head)
		if end == nil {
			return head
		}
		pre.Next = nil
		end.Next = head
		head = end
	}
	return head
}
func ListLen(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
func LastNode(head *ListNode) (pre, last *ListNode) {
	if head != nil && head.Next != nil && head.Next.Next == nil {
		return head, head.Next
	}
	if head.Next == nil {
		return head, nil
	}
	return LastNode(head.Next)
}

func rotateRight0(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
