package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
}

func isValid(s string) bool {
	st := []string{}
	for _, v := range s {
		char := string(v)
		if char == "(" || char == "{" || char == "[" {
			st = Push(st, leftChar(char))
		} else {
			if len(st) > 0 && char == st[0] {
				st = Pop(st)
			} else {
				return false
			}
		}
	}
	return len(st) == 0
}
func Pop(nums []string) []string {
	return nums[1:]
}

func Push(nums []string, num string) []string {
	temp := []string{num}
	temp = append(temp, nums...)
	return temp
}
func leftChar(char string) string {
	if char == "{" {
		return "}"
	}
	if char == "(" {
		return ")"
	}
	return "]"
}