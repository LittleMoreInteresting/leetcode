package p167

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
