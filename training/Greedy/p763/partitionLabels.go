package main

import "fmt"

func partitionLabels(s string) []int {

	labels := make(map[rune]int)
	for i, v := range s {
		labels[v] = i
	}
	result := []int{}
	left := 0
	right := labels[rune(s[0])]
	for i, v := range s {
		if i > right {
			result = append(result, right-left+1)
			left = i
		}
		right = Max(right, labels[v])
	}
	result = append(result, right-left+1)
	return result
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	result := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Println(result)
}
