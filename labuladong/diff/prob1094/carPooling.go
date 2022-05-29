package main

import (
	"fmt"
	"leetcode/labuladong/diff"
)
func carPooling(trips [][]int, capacity int) bool {
	numbs := make([]int, 1000)
	df := diff.New(numbs)
	for _, v := range trips {
		inc := v[0]
		i := v[1]
		j := v[2] -1 
		df.Increment(i,j,inc)
	}
	result := df.Result()
	for _, v := range result {
		if capacity < v {
			return false
		}
	}
	return true
}

func main() {
	//trips := [][]int{{2,1,5},{3,3,7}}
	//capacity := 4 
	// false
	trips := [][]int{{2,1,5},{3,3,7}}
	capacity := 5
	fmt.Println(carPooling(trips,capacity))
}