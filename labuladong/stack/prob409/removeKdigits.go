package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("1432219",3))
	fmt.Println(removeKdigits("10200",1))
	fmt.Println(removeKdigits("10",2))
}

func removeKdigits(num string, k int) string {
	st := []rune{}
	for _, v := range num {
		for len(st) > 0 && v < st[0] && k > 0 {
			st = pop(st)
			k--
		}
		if len(st) == 0 && string(v) == "0" {
			continue
		}
		st = push(st, v)
	}

	for k > 0 && len(st) > 0 {
		st = pop(st)
		k--
	}

	if len(st) == 0 {
		return "0"
	}
	ret := ""
	for i := len(st) - 1; i >= 0; i-- {
		ret = ret + string(st[i])
	}
	return ret
}

func pop(nums []rune) []rune {
	return nums[1:]
}

func push(nums []rune, num rune) []rune {
	temp := []rune{num}
	temp = append(temp, nums...)
	return temp
}