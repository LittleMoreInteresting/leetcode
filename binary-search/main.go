package binary_search

func BinarySearch(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	left := 0
	right := l - 1
	for {
		if left > right {
			return -1
		}
		mid := (left + right) / 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			left = 0
			right = mid - 1
			continue
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
}
