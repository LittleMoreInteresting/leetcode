package main

import "fmt"

func main() {
	fmt.Println(extent("helo", "heeellooo"))
}

func expressiveWords(s string, words []string) int {
	count := 0
	for _, w := range words {
		if extent(w, s) {
			count++
		}
	}
	return count
}

func extent(w, s string) bool {
	m, n := len(w), len(s)
	i, j := 0, 0
	for i < m && j < n {
		if w[i] != s[j] {
			return false
		}
		char := w[i]
		cntI := 0
		for i < m && w[i] == char {
			cntI++
			i++
		}
		cntJ := 0
		for j < n && s[j] == char {
			cntJ++
			j++
		}
		if cntI > cntJ || cntI < cntJ && cntJ < 3 {
			return false
		}
	}
	return i == m && j == n
}
