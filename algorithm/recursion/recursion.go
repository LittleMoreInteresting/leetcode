package recursion

func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursion(n-1)
}

//尾部递归
func RecursionTail(m, n int) int {
	if m == 1 {
		return n
	}
	return RecursionTail(m-1, m*n)
}

func FiboTail(m, n, q int) int {
	if m == 0 {
		return n
	}
	return FiboTail(m-1, q, n+q)
}

//二分查找 随机
func BinarySearch(array []int, target int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	middleVal := array[mid]
	if middleVal == target {
		return mid
	}
	if middleVal > target {
		return BinarySearch(array, target, 0, mid-1)
	}
	return BinarySearch(array, target, mid+1, right)
}

func BinarySearchRight(array []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	middleVal := array[mid]
	if middleVal == target {
		nextLeft := BinarySearch(array, target, mid+1, right)
		if nextLeft != -1 {
			return nextLeft
		}
		return mid
	}
	if middleVal > target {
		return BinarySearch(array, target, 0, mid-1)
	}
	return BinarySearch(array, target, mid+1, right)
}
