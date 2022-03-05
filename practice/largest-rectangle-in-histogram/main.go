package main

import "fmt"

func largestRectangleArea(heights []int) int {
	var largestArea int
	hlen := len(heights)
	for i := 0; i < hlen; i++ {
		h := heights[i] //高度
		left := i - 1
		right := i + 1
		//左边界
		for ; left >= 0; left-- {
			if heights[left] < h {
				break
			}
		}
		for ; right < hlen; right++ {
			if heights[right] < h {
				break
			}
		}
		area := h * (right - left - 1)
		if area > largestArea {
			largestArea = area
		}
		fmt.Printf("%d*(%d-%d-1)=%d\n", h, right, left, area)
	}

	return largestArea
}

func main() {
	heights := []int{1, 1}
	//heights := []int{2}
	//heights := []int{2}
	area := largestRectangleArea(heights)
	fmt.Println(area)
}
