package p33

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left < right && nums[left] == nums[left+1] {
		left++
	}
	for left < right && nums[right] == nums[right-1] {
		right--
	}

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return true
		}
		if nums[mid] >= nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
