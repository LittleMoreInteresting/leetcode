package p92

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for i := 0; i < m-1; i++ {
		p = p.Next
	}
	cur := p.Next
	var last *ListNode
	for i := m; i < n; i++ {
		last = cur.Next
		cur.Next = last.Next
		last.Next = p.Next
		p.Next = last
	}
	return dummy.Next
}
