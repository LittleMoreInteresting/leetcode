package main

import "fmt"

func main() {
	head := &ListNode{2, &ListNode{1, &ListNode{5, nil}}}
	fmt.Println(nextLargerNodes(head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	nums := []int{}
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	fmt.Println(nums)
	st := []int{}
	size := 0
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for size > 0 && st[size-1] <= nums[i] {
			st = st[:size-1]
			size--
		}
		if size == 0 {
			res[i] = 0
		} else {
			res[i] = st[size-1]
		}
		st = append(st, nums[i])
		size++
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
