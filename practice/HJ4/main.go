package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	n := len(str)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print(string(str[i]))
		if i > 7 && i%7 == 0 {
			fmt.Println()
		}
	}
	if n%8 != 0 {
		for i := n % 7; i < 8; i++ {
			fmt.Print(0)
		}
		fmt.Println()
	}

}
