package main

import "fmt"

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"abc"}, []string{"a", "b", "c"}))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	idx1, idx2 := 0, 0
	i, j := 0, 0
	for idx1 < len(word1) && idx2 < len(word2) {
		if word1[idx1][i] != word2[idx2][j] {
			return false
		}
		i++
		if i == len(word1[idx1]) {
			idx1++
			i = 0
		}
		j++
		if j == len(word2[idx2]) {
			idx2++
			j = 0
		}
	}
	return idx1 == len(word1) && idx2 == len(word2)
}
