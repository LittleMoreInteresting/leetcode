package main

import "fmt"

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

var res []string

func letterCasePermutation(s string) []string {
	res = []string{}
	letterCaseP(s, 0)
	return res
}

func letterCaseP(s string, start int) {
	for start < len(s) && s[start] >= '0' && s[start] <= '9' {
		start++
	}
	if start >= len(s) {
		res = append(res, s)
		return
	}
	s2 := []byte(s)
	charLen := 'a' - 'A'
	if s2[start] >= 'a' && s2[start] <= 'z' {
		s2[start] = s2[start] - byte(charLen)
	} else {
		s2[start] = s2[start] + byte(charLen)
	}
	letterCaseP(s, start+1)
	letterCaseP(string(s2), start+1)
}
