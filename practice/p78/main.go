package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

var res = [][]int{}
var track = []int{}

func subsets(nums []int) [][]int {
	res = [][]int{}
	track = []int{}
	backtrack(nums, 0)
	return res
}

func backtrack(nums []int, start int) {
	trackCopy := make([]int, len(track))
	copy(trackCopy, track)
	res = append(res, trackCopy)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1)
		track = track[:len(track)-1]
	}
}
