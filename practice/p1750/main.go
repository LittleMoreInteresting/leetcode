package main

import "fmt"

func main() {
	fmt.Println(minimumLength("cabaabac"))
}

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		c := s[left]
		if s[left] == s[right] {
			left++
			for left < right && s[left] == c {
				left++
			}
			right--
			for left < right && s[right] == c {
				right--
			}

		} else {
			break
		}
	}
	fmt.Println(left, right)
	return right - left + 1
}
