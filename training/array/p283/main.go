package p283

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for j := n; j < len(nums); j++ {
		nums[j] = 0
	}
}
