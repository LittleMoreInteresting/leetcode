package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	li := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		li = append(li, charEncode(s[i]))
	}
	fmt.Println(string(li))

	input.Scan()
	s = input.Text()
	li = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		li = append(li, charDecode(s[i]))
	}
	fmt.Println(string(li))
}

func charEncode(char byte) byte {
	if char == 'Z' {
		return 'a'
	} else if char == 'z' {
		return 'A'
	} else if char == '9' {
		return '0'
	} else if char >= 'a' && char < 'z' {
		return char - 32 + 1
	} else if char >= 'A' && char < 'Z' {
		return char + 32 + 1
	} else if char >= '0' && char < '9' {
		return char + 1
	} else {
		return char
	}
}

func charDecode(char byte) byte {
	if char == 'a' {
		return 'Z'
	} else if char == 'A' {
		return 'z'
	} else if char == '0' {
		return '9'
	} else if char > 'a' && char <= 'z' {
		return char - 32 - 1
	} else if char > 'A' && char <= 'Z' {
		return char + 32 - 1
	} else if char > '0' && char <= '9' {
		return char - 1
	} else {
		return char
	}
}
