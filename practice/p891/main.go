package main

import "sort"

func main() {

}

func sumSubseqWidths(nums []int) int {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	res, x, y := 0, nums[0], 2
	for _, num := range nums[1:] {
		res = (res + num*(y-1) - x) % mod
		x = (x*2 + num) % mod
		y = y * 2 % mod
	}
	return (res + mod) % mod
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/sum-of-subsequence-widths/solutions/1976655/zi-xu-lie-kuan-du-zhi-he-by-leetcode-sol-649q/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
