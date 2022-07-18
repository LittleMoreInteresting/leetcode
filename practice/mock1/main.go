// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	c := 0
	nums := []int{1, 2, 3, 4, 1, 2, 3, 4}

	var i int = 0
	for i < len(nums) {
		tag := findTag(nums, nums[i], i+1)
		if tag != -1 {
			nums = append(nums[0:tag], nums[tag+1:]...)
			nums = append(nums[0:i], nums[i+1:]...)
			c++
		} else {
			i++
		}
	}
	fmt.Println(nums)
	fmt.Println(c)
}

func findTag(nums []int, target, s int) int {
	for i := s; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
