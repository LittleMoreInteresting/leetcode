package main

import "fmt"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}

func move(target, step int) int {
	if target == 0 {
		return step
	}
	step++
	left := move(target-step, step)
	right := move(target+step, step)
	return Min(left, right)
}
func Min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
func main() {
	fmt.Println(reachNumber(2))
}
