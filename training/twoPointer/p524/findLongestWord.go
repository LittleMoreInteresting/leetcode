package main

import "fmt"

func findLongestWord(s string, dictionary []string) string {
	okWorld := ""
	for _, word := range dictionary {
		if len(word) < len(okWorld) {
			continue
		}
		si, di, okCount := 0, 0, 0
		for di < len(word) && si < len(s) {
			if s[si] == word[di] {
				si++
				di++
				okCount++
			} else {
				si++
			}
		}
		if okCount == len(word) {
			if len(word) > len(okWorld) ||
				(len(word) == len(okWorld) && word < okWorld) {
				okWorld = word
			}
		}
	}
	return okWorld
}

func main() {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	fmt.Println(findLongestWord(s, d))
	s = "abpcplea"
	d = []string{"a", "b", "c"}
	fmt.Println(findLongestWord(s, d))
}
