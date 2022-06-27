package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	res := make(map[int]int)
	sortKeys := []int{}
	for i := 0; i < n; i++ {
		var k, v int
		fmt.Scan(&k, &v)
		if _, ok := res[k]; ok {
			res[k] += v
		} else {
			res[k] = v
			sortKeys = append(sortKeys, k)
		}
	}
	sort.Slice(sortKeys, func(i, j int) bool {
		if sortKeys[i] < sortKeys[j] {
			return true
		}
		return false
	})
	for _, value := range sortKeys {
		fmt.Printf("%d %d\n", value, res[value])
	}
}
