package main

import (
	"fmt"
	df "leetcode/labuladong/diff"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	numbs := make([]int, n)
	d := df.New(numbs)
	for _, v := range bookings {
		d.Increment(v[0]-1, v[1]-1, v[2])
	}
	return d.Result()
}

func main() {
	//updates := [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}
	updates := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}

	res := corpFlightBookings(updates, 5)
	fmt.Println(res)
}