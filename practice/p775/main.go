package main

import (
	"fmt"
)

func main() {
	rs := isIdealPermutation([]int{1, 0, 3})
	fmt.Println(rs)
	rs = isIdealPermutation([]int{1, 2, 0})
	fmt.Println(rs)
}
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minNum := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minNum {
			return false
		}
		minNum = min(minNum, nums[i])
	}
	return true

	/*作者：力扣官方题解
	链接：https://leetcode.cn/problems/global-and-local-inversions/solutions/1971259/quan-ju-dao-zhi-yu-ju-bu-dao-zhi-by-leet-bmjp/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
