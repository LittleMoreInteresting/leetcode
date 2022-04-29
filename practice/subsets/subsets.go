package main

import (
	"fmt"
	"sort"
)

var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	backtrack([]int{}, nums, 0)
	return res
}
func backtrack(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtrack(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	num := []int{1,2,3}
	ret := subsets(num)
	fmt.Printf("%v \n",ret)
	num = []int{1}
	ret = subsets(num)
	fmt.Printf("%v \n",ret)
}