package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	mem := make(map[int32]bool)
	for _, c := range str {
		mem[c] = true
	}
}
