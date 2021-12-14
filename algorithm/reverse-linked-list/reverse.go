package reverse_linked_list

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

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		/*next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next*/
		cur, cur.Next, prev = cur.Next, prev, cur

	}
	return prev
}
