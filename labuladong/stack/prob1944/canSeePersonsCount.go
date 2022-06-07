package main

import "fmt"

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}

func canSeePersonsCount(heights []int) []int {
	l := len(heights)
	res := make([]int, l)
	if l == 0 {
		return res
	}
	st := []int{}

	size := 0
	for i := l - 1; i >= 0; i-- {
		count := 0
		for size > 0 && st[size-1] < heights[i] {
			st = st[:size-1]
			size--
			count++
		}
		if size == 0 {
			res[i] = count
		} else {
			res[i] = count + 1
		}
		st = append(st, heights[i])
		size++
	}
	return res
}