package main

import (
	"fmt"
	binarysearch "leetcode/labuladong/binarySearch"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8
	ret := []int{binarysearch.LeftBound(nums, target),binarysearch.RightBound(nums, target)}
	fmt.Println(ret)
}