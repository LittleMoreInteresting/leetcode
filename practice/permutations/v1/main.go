package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	track, used := []int{}, make([]bool, len(nums))
	backtrack(nums, track, used)
	return res
}

func backtrack(nums, track []int, used []bool) {
	if len(nums) == len(track) {
		cp := make([]int, len(track))
		copy(cp, track) //消除切片底层公用数据影响
		res = append(res, cp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}
		track = append(track, nums[i])
		//fmt.Println(track)
		//fmt.Println(res)
		used[i] = true
		backtrack(nums, track, used)

		track = track[0 : len(track)-1]
		used[i] = false
	}
}

func main() {
	//nums := []int{1,2,3};
	//nums := []int{0,1};
	nums := []int{5, 4, 6, 2}
	i := permute(nums)
	fmt.Println(i)
}
