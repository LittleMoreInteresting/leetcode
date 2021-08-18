package main

func moveZeroes(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{1, 3, 4, 2}
	//nums := []int{0,1,0,3,12}
	moveZeroes(nums)
}