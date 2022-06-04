package stack

func Pop(nums []int) []int {
	return nums[1:]
}

func Push(nums []int, num int) []int {
	temp := []int{num}
	temp = append(temp, nums...)
	return temp
}