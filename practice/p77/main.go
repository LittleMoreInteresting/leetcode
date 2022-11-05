package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

var res = [][]int{}
var track = []int{}

func combine(n int, k int) [][]int {
	res = [][]int{}
	track = []int{}
	backtrack(n, 1, k)
	return res
}
func backtrack(nums int, start, k int) {
	if k == len(track) {
		trackCopy := make([]int, len(track))
		copy(trackCopy, track)
		res = append(res, trackCopy)
	}

	for i := start; i <= nums; i++ {
		track = append(track, i)
		backtrack(nums, i+1, k)
		track = track[:len(track)-1]
	}
}
