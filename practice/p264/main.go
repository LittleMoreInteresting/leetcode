package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	pVal2, pVal3, pVal5 := 1, 1, 1
	p := 1
	result := make([]int, n+1)
	for p <= n {
		minVal := min(pVal2, pVal3, pVal5)
		result[p] = minVal
		fmt.Println(minVal)
		p++
		if minVal == pVal2 {
			pVal2 = 2 * result[p2]
			p2++
		}
		if minVal == pVal3 {
			pVal3 = 3 * result[p3]
			p3++
		}
		if minVal == pVal5 {
			pVal5 = 5 * result[p5]
			p5++
		}
	}
	return result[n]
}

func min(a ...int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
