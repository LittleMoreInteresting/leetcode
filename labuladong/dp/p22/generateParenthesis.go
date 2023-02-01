package main

import "fmt"

var res []string

func main() {
	r3 := generateParenthesis(3)
	fmt.Println(r3)
	r2 := generateParenthesis(2)
	fmt.Println(r2)
	r1 := generateParenthesis(1)
	fmt.Println(r1)
}
func generateParenthesis(n int) []string {
	res = []string{}
	dp(n, -1, "(")
	return res
}

func dp(n, tag int, s string) {
	if n*2 == len(s) || tag > 0 || -tag > n {
		if tag == 0 {
			res = append(res, s)
		}
		fmt.Println(s, tag, res)
		return
	}
	if -tag < n {
		dp(n, tag-1, s+"(")
	}
	if tag < 0 {
		dp(n, tag+1, s+")")
	}
}
