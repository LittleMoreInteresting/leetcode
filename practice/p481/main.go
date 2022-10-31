package main

import "fmt"

func main() {
	fmt.Println(magicalString(4))
	fmt.Println(magicalString1(4))
}

func magicalString(n int) int {
	if n < 4 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 1, 2, 2
	cur := 1
	left := 2
	right := 3
	res := 1
	for right < n {

		if arr[left] == 1 {
			arr[right] = cur
			if cur == 1 && right < n {
				res++
			}
			right++
		} else {
			arr[right] = cur
			if cur == 1 && right < n {
				res++
			}
			right++
			arr[right] = cur
			if cur == 1 && right < n {
				res++
			}
			right++
		}
		cur = 3 - cur
		left++
	}
	fmt.Println(arr)
	return res
}

func magicalString1(n int) int {
	if n < 4 {
		return 1
	}
	s := make([]byte, n)
	copy(s, "122")
	res := 1
	i, j := 2, 3
	for j < n {
		size := s[i] - '0'
		num := 3 - (s[j-1] - '0')
		for size > 0 && j < n {
			s[j] = '0' + num
			if num == 1 {
				res++
			}
			j++
			size--
		}
		i++
	}
	return res
}
