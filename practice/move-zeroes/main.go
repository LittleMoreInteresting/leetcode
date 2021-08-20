package main

import "fmt"

/*func moveZeroes(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}*/

func moveZeros2(nums []int) {
	n := len(nums)
	zeroNum := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zeroNum++
		} else {
			index := i - zeroNum
			nums[index] = nums[i]
			if zeroNum != 0 {
				nums[i] = 0
			}
		}
	}
}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	//nums := []int{1, 3, 4, 2}
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("%v", nums)
	nums2 := []int{0, 1, 0, 3, 12}
	moveZeros2(nums2)
	fmt.Printf("%v", nums2)
}
