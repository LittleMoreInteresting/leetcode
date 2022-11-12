package p2

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p1, p2, pSum := l1, l2, dummy
	toHeight := 0
	for p1 != nil && p2 != nil {
		sum := p1.Val + p2.Val + toHeight
		toHeight = sum / 10
		val := sum % 10
		pSum.Next = &ListNode{Val: val}
		pSum = pSum.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		sum := p1.Val + toHeight
		toHeight = sum / 10
		val := sum % 10
		pSum.Next = &ListNode{Val: val}
		pSum = pSum.Next
		p1 = p1.Next
	}
	for p2 != nil {
		sum := p2.Val + toHeight
		toHeight = sum / 10
		val := sum % 10
		pSum.Next = &ListNode{Val: val}
		pSum = pSum.Next
		p2 = p2.Next
	}
	if toHeight > 0 {
		pSum.Next = &ListNode{Val: toHeight}
	}
	return dummy.Next
}
