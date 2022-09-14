package main

import "fmt"

func minWindow(s string, t string) string {
	window, taget := make(map[uint8]int), make(map[uint8]int)
	for i := range t {
		taget[t[i]]++
	}
	left, right, okCount := 0, 0, 0
	sLen, tLen := len(s), len(taget)
	start, end := 0, sLen+1
	for right < sLen {
		c := s[right]
		right++ //
		if _, ok := taget[c]; ok {
			window[c]++
			if window[c] == taget[c] {
				okCount++
			}
		}
		for okCount == tLen { //
			if right-left < end-start {
				start, end = left, right
			}
			r := s[left]
			left++
			if _, ok := taget[r]; ok {
				if window[r] == taget[r] {
					okCount--
				}
				window[r]--
			}
		}

	}
	if end == sLen+1 {
		return ""
	}
	return s[start:end]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
