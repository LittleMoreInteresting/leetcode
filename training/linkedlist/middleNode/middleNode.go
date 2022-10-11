package middleNode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
寻找链表中间节点
*/
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
