package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[string]int{}, map[string]int{}

	for _, v := range s1 {
		chr := string(v)
		if _, ok := need[chr]; !ok {
			need[chr] = 1
			continue
		}
		need[chr]++
	}

	left, right, valid := 0, 0, 0

	for right < len(s2) {
		c := string(s2[right])
		right++ //窗口右移
		//窗口内数据更新
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
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := string(s2[left])
			left++

			//窗口内数据更新
			if v, ok := need[d]; ok {
				if v == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("a", "aa"))
}