package p21

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	start := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			start.Next = p2
			p2 = p2.Next
		} else {
			start.Next = p1
			p1 = p1.Next
		}
		start = start.Next
	}
	if p1 != nil {
		start.Next = p1
	}
	if p2 != nil {
		start.Next = p2
	}
	return dummy.Next
}
