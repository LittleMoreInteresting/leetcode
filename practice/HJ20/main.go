package main

import (
	"fmt"
	"strings"
)

func main() {
	var str []string
	for {
		var temp string
		fmt.Scan(&temp)
		if len(temp) == 0 {
			break
		}
		str = append(str, temp)
	}

	for _, v := range str {
		fmt.Println(strCheck(v))
	}
}

func strCheck(str string) string {
	if len(str) < 8 {
		return "NG"
	}
	var mode int
	for _, v := range str {
		mode |= charToMode(v)
	}
	level := 0
	for i := 0; i < 4; i++ {
		if mode&1 == 1 {
			level++
		}
		mode >>= 1
	}
	if level < 3 {
		return "NG"
	}
	if childCheck(str) {
		return "NG"
	}
	return "OK"
}

func charToMode(n int32) int {
	if n >= 48 && n <= 57 {
		return 1
	}
	if n >= 65 && n <= 90 {
		return 2
	}
	if n >= 97 && n <= 120 {
		return 4
	}
	return 8
}

func childCheck(str string) bool {
	l := len(str)
	for i := 3; i < l/2; i++ {
		for j := 0; j < l-i; j++ {
			ch := str[j : i+j]
			left := str[i+j:]

			if strings.Contains(left, ch) {
				return true
			}
		}
	}
	return false
}
