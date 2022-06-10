package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse","ros"))
	fmt.Println(minDistance("intention","execution"))
}

var mem map[string]int

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	mem = make(map[string]int, m)
	return dp(word1, m-1, word2, n-1)
}

func dp(word1 string, i int, word2 string, j int) int {
	//base case
	if i == -1 {
		return j+1
	}
	if j == -1 {
		return i+1
	}
	key := fmt.Sprintf("%d-%d", i, j)
	if v,ok := mem[key]; ok {
		return v
	}
	//状态转移
	if word1[i] == word2[j] {
		return dp(word1, i-1, word2, j-1)
	}
	mem[key] = min(
		dp(word1, i, word2, j-1)+1,
		dp(word1, i-1, word2, j)+1,
		dp(word1, i-1, word2, j-1)+1,
	)
	return mem[key]
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}