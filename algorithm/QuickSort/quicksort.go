package main

import "fmt"

func Partition0(nums *[]int, low, height int) int {
	i, j, pivot := low, height, (*nums)[low]
	for i < j {
		for i < j && (*nums)[j] > pivot {
			j--
		}
		if i < j {
			Swap(nums, i, j)
			i++
		}
		for i < j && (*nums)[i] <= pivot {
			i++
		}
		if i < j {
			Swap(nums, i, j)
			j--
		}
	}
	return i
}
func Partition(nums *[]int, low, height int) int {
	i, j, pivot := low, height, (*nums)[low]
	for i < j {
		for i <= j && (*nums)[j] > pivot {
			j--
		}
		for i <= j && (*nums)[i] < pivot {
			i++
		}
		if i < j {
			Swap(nums, i, j)
			i++
			j--
		}
	}
	if (*nums)[i] > pivot {
		Swap(nums, i-1, low)
		return i - 1
	}
	Swap(nums, i, low)
	return i
}

func Swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func QuickSort(nums *[]int, low, height int) {
	if low < height {
		mid := Partition(nums, low, height)
		QuickSort(nums, low, mid-1)
		QuickSort(nums, mid+1, height)
	}
}
func main() {
	nums := []int{2, 4, 2, 5, 762, 141}
	QuickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{}
	QuickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{5, 2}
	QuickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{2, 2, 2, 2, 2, 2, 2, 2}
	QuickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
}
