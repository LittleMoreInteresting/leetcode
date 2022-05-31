package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	right, left := 0, 0
	window := map[byte]int{}
	result := 0
	for right < len(s) {
		// move right
		c := s[right]
		right++
		// updage window
		if _, ok := window[c]; !ok {
			window[c] = 1
		} else {
			window[c]++
		}
		// move left
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if right-left > result {
			result = right - left
		}
	}
	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}