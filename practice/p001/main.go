package main

import (
	"fmt"
	"sort"
)

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

答案中不可以包含重复的三元组。

*/

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := zeroSum(nums, 0)
	fmt.Println(res)
}

func zeroSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n; i++ {
		twoRes := twoSum(nums, i+1, target-nums[i])

		for _, t := range twoRes {
			t = append(t, nums[i])
			res = append(res, t)
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	low, height := start, len(nums)-1
	res := [][]int{}
	for low < height {
		sum := nums[low] + nums[height]
		lnum, hnum := nums[low], nums[height]
		fmt.Println(nums, sum, low, height, target)
		if sum == target {
			res = append(res, []int{nums[low], nums[height]})
			for low < height && nums[height] == hnum {
				height--
			}
			for low < height && nums[low] == lnum {
				low++
			}
		} else if sum > target {
			height--
		} else if sum < target {
			low++
		}
	}
	return res
}
