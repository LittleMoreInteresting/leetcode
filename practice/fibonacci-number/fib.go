package main

import (
	"fmt"
	"time"
)

// 傻递归
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 备忘录
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp_i_1, dp_i_2 := 1, 0
	for i := 2; i <= n; i++ {
		dp_i := dp_i_1 + dp_i_2
		//记录
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}

func main() {
	bg := time.Now()
	//ret := fib(100)
	ret := fib2(100)
	fmt.Printf("%d,%v", ret, time.Since(bg).Milliseconds())
}
