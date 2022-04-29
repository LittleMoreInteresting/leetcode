package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	track, used := []int{}, make([]bool, len(nums))
	backtrack(nums, track, used)
	return res
}

func backtrack(nums []int, track []int, used []bool) {
	if len(nums) == len(track) {
		tmp := make([]int,len(track))
		copy(tmp,track)
		res = append(res,tmp)
		return 
	}
	for i := 0;i<len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i]);
		used[i] = true
		backtrack(nums,track,used)
		track = track[:len(track)-1]
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