package main

import "fmt"

func GenMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = i*n + j
		}
	}
	return m
}

func ShowMatrix(num [][]int) {
	for i := range num {
		for _, v := range num[i] {
			fmt.Printf("| %d ", v)
		}
		fmt.Println("|")
	}
	fmt.Println("-------------")
}
func main() {
	N := 4
	m := GenMatrix(N)
	ShowMatrix(m)
	n := Rotation(m, N)
	ShowMatrix(n)
	q := Rotation1(m, N)
	ShowMatrix(q)
}

// y[i][j]=x[j][n-1-i]; 逆时针
func Rotation(num [][]int, n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = num[j][n-1-i]
		}
	}
	return m
}

// 顺时针
func Rotation1(num [][]int, n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = num[n-1-j][i]
		}
	}
	return m
}
