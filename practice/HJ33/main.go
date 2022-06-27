package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	a := input.Text()
	input.Scan()
	b, _ := strconv.Atoi(input.Text())
	li := strings.Split(a, ".")
	req1 := 0
	for i := 0; i < len(li); i++ {
		n, _ := strconv.Atoi(li[i])
		req1 = req1*256 + n
	}
	req2 := ""
	for i := 0; i < len(li); i++ {
		if i == 0 {
			req2 = strconv.Itoa(b % 256)
		} else {
			req2 = strconv.Itoa(b%256) + "." + req2
		}
		b = b / 256
	}

	fmt.Println(req1)
	fmt.Println(req2)
}
