package main

import "fmt"

func main() {
	fmt.Println(interpret("G()(al)"))
	fmt.Println(interpret("G()()()()(al)"))
}

func interpret(command string) string {
	left, right := 0, 0
	res := []byte{}
	for right < len(command) {
		right++
		str := command[left:right]
		if str == "G" {
			res = append(res, 'G')
			left = right
			continue
		}
		if str == "()" {
			res = append(res, 'o')
			left = right
			continue
		}
		if str == "(al)" {
			res = append(res, 'a', 'l')
			left = right
			continue
		}
	}
	return string(res)
}
