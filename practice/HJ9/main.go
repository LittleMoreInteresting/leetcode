package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	mem := make(map[int]bool)
	res := 0
	m := 10
	for num > 0 {
		v := num % m

		if _, ok := mem[v]; !ok {
			res = res*m + v
			mem[v] = true
		}
		num = num / m
	}
	fmt.Print(res)
}
