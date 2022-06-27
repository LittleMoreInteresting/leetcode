package main

import (
	"fmt"
	"sort"
)

func main() {
	var line int
	fmt.Scan(&line)
	res := []int{}
	mem := make(map[int]bool)
	for i := 0; i < line; i++ {
		var num int
		fmt.Scan(&num)
		if v := mem[num]; v {
			continue
		}
		mem[num] = true
		res = append(res, num)
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i] < res[j] {
			return true
		}
		return false
	})
	for _, v := range res {
		fmt.Println(v)
	}
}
