package main

import (
	"fmt"
)

func main() {
	var num float32
	fmt.Scan(&num)
	fmt.Println(int32(num + 0.5))
}
