package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1000000000
	for left <= right {
		mid := left + (right-left)/2
		if needHours(piles, mid) <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 速度 => 时间关系
func needHours(piles []int, speed int) int {
	hours := 0
	for _, v := range piles {
		hours += v / speed
		if v%speed != 0 {
			hours++
		}
	}
	return hours
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	result := minEatingSpeed(piles, h)
	fmt.Println(result)
	piles = []int{30,11,23,4,20}
	h = 5
	result = minEatingSpeed(piles, h)
	fmt.Println(result)
	h = 6
	result = minEatingSpeed(piles, h)
	fmt.Println(result)
}