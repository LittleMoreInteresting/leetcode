package binarysearch

// 基本的二分搜索
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 左边界二分查找
func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1 // 继续向左侧区间搜索
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 右边界二分查找
func RightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1 // 继续向右侧区间搜索
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	if right <0 || nums[right] != target {
		return -1
	}
	return right
}