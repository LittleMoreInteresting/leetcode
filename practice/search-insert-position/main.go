package search_insert_position

func searchInsert(nums []int, target int) int {
	length := len(nums)
	l := 0
	r := length - 1
	for {
		if l > r {
			break
		}
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			l = 0
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return l
}
