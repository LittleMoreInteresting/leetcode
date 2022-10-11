package p19

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

func findNthNode(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	target := findNthNode(dummy, n+1)
	target.Next = target.Next.Next
	return dummy.Next
}
