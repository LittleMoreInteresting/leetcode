package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	c, _ := reader.ReadString('\n')
	fmt.Println(strCount([]byte(str), []byte(c)))
}

func strCount(str, c []byte) int {
	if len(c) == 0 || len(str) == 0 {
		return 0
	}
	count := 0
	for _, v := range str {
		if charToUpper(v) == charToUpper(c[0]) {
			count++
		}
	}
	return count
}

func charToUpper(x byte) byte {
	if x >= 65 && x <= 90 {
		return x + 32
	}
	return x
}
