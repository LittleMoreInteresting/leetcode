package main

import "fmt"

// 速度，天数 关系
func needDays(weights []int, x int) int {
	days := 0
	sum := 0
	for _, v := range weights {
		if (sum + v) > x {
			days++
			sum = v
		} else {
			sum += v
		}
	}
	if (sum>0) {
		days++
	}
	return days
}

func shipWithinDays(weights []int, days int) int {
	left :=0
	right := 0
	for _, v := range weights {
		left = max(left,v)
		right += v
	}
	for left <= right {
		mid := left + (right - left)/2
		if (needDays(weights,mid) <= days){
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return left
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(shipWithinDays(weights, 5)) //15
	weights = []int{3,2,2,4,1,4}
	fmt.Println(shipWithinDays(weights, 3)) //6
	weights = []int{1,2,3,1,1}
	fmt.Println(shipWithinDays(weights, 4)) //3
}