package p2032

func twoOutOfThree(nums1, nums2, nums3 []int) []int {
	mark := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mark[x] |= 1 << i
		}
	}
	ans := []int{}
	for x, m := range mark {
		if m&(m-1) > 0 {
			ans = append(ans, x)
		}
	}
	return ans
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/two-out-of-three/solutions/2034884/zhi-shao-zai-liang-ge-shu-zu-zhong-chu-x-5131/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
