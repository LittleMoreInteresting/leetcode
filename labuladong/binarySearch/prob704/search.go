package main

import (
	"fmt"
	binarysearch "leetcode/labuladong/binarySearch"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	ret := binarysearch.BinarySearch(nums, target)
	fmt.Println(ret)
}