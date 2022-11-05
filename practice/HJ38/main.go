package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	coins := make([]int, t)
	set := map[int]struct{}{0: struct{}{}}
	keys := []int{0}
	for i := 0; i < t; i++ {
		var w int
		fmt.Scan(&w)
		coins[i] = w
		for _, k := range keys {
			if _, ok := set[k+w]; !ok {
				set[k+w] = struct{}{}
				keys = append(keys, k+w)
			}
		}
	}
	for i := 0; i < t; i++ {
		var num int
		fmt.Scan(&num)
		for j := 0; j < num-1; j++ {
			w := coins[i]
			for _, k := range keys {
				if _, ok := set[k+w]; !ok {
					set[k+w] = struct{}{}
					keys = append(keys, k+w)
				}
			}
		}
	}
	fmt.Println(len(keys))
}
