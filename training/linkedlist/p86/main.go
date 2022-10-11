package p86

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

func partition(head *ListNode, x int) *ListNode {
	bg, lg := &ListNode{}, &ListNode{}
	p := head
	pb, pl := bg, lg
	for p != nil {
		if p.Val >= x {
			pb.Next = p
			pb = pb.Next
		} else {
			pl.Next = p
			pl = pl.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	pl.Next = bg.Next
	return lg.Next
}
