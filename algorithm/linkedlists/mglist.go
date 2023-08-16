package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	h := &ListNode{}
	p := h
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val > pHead2.Val {
			p.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			p.Next = pHead2
			pHead2 = pHead2.Next
		}
		p = p.Next
	}
	if pHead1 != nil {
		p.Next = pHead1
	}
	if pHead2 != nil {
		p.Next = pHead2
	}
	return h.Next
}
