package main

import "fmt"

func findAnagrams(s string, p string) []int {
	need, window := map[byte]int{}, map[byte]int{}
	for _, v := range p {
		chr := byte(v)
		if _, ok := need[chr]; !ok {
			need[chr] = 1
			continue
		}
		need[chr]++
	}
	result := []int{}
	var right, left, valid int = 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		//窗口操作
		if _, ok := need[c]; ok {
			if _, ok := window[c]; !ok {
				window[c] = 1
			} else {
				window[c]++
			}
			if need[c] == window[c] {
				valid++
			}
		}
		//判断左侧窗口是否需要更新
		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}
			b := s[left]
			left++
			if _, ok := need[b]; ok {
				if need[b] == window[b] {
					valid--
				}
				window[b]--
			}
		}

	}
	return result

}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}