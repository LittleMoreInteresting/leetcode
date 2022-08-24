package main

import "fmt"

func Merge(A *[]int, low, mid, height int) {
	B := make([]int, height-low+1)
	i, j, k := low, mid+1, 0
	for i <= mid && j <= height {
		if (*A)[i] <= (*A)[j] {
			B[k] = (*A)[i]
			i++
		} else {
			B[k] = (*A)[j]
			j++
		}
		k++
	}
	for i <= mid {
		B[k] = (*A)[i]
		i++
		k++
	}
	for j <= height {
		B[k] = (*A)[j]
		j++
		k++
	}
	i, k = low, 0
	for i <= height {
		(*A)[i] = B[k]
		i++
		k++
	}
}

//合并排序
func MergeSort(A *[]int, low, height int) {
	if low < height {
		mid := low + (height-low)/2
		MergeSort(A, low, mid)
		MergeSort(A, mid+1, height)
		Merge(A, low, mid, height)
	}
}

func main() {
	nums := []int{2, 4, 2, 5, 762, 141}
	MergeSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{}
	MergeSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{2, 2, 3, 2, 3, 2, 2, 2}
	MergeSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
}
