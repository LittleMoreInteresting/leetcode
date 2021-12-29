package linked_list_cycle

import "fmt"

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

//快慢指针
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 借助set 存储节点
func hasCycle_V1(head *ListNode) bool {
	nodePool := map[string]int{}
	current := head
	for current != nil {
		key := fmt.Sprintf("k-%v", current)
		if _, ok := nodePool[key]; ok {
			return true
		}
		nodePool[key] = 1
		current = current.Next
	}
	return false
}
