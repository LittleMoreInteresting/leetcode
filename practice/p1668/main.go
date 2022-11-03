package main

import "fmt"

func main() {
	fmt.Println(maxRepeating("ababc", "ac"))
	fmt.Println(maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))
	fmt.Println(maxRepeating2("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))
	fmt.Println(maxRepeating3("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))
}

func maxRepeating(sequence string, word string) int {
	sl, wl := len(sequence), len(word)
	repeat := 0

	for i := 0; i < sl; i++ {
		k := i
		rep := 0
		j := 0
		for k < sl && sequence[k] == word[j] {
			j++
			k++
			if j == wl {
				j = 0
				rep++
			}
		}
		if rep > repeat {
			repeat = rep
		}

	}
	return repeat
}

func maxRepeating2(sequence, word string) (ans int) {
	n, m := len(sequence), len(word)
	if n < m {
		return
	}
	f := make([]int, n)
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				f[i] = 1
			} else {
				f[i] = f[i-m] + 1
			}
			if f[i] > ans {
				ans = f[i]
			}
		}
	}
	return
}

func maxRepeating3(sequence, word string) (ans int) {
	n, m := len(sequence), len(word)
	if n < m {
		return
	}
	f := make([]int, n)
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				f[i] = 1
			} else {
				f[i] = f[i-m] + 1
			}
			if f[i] > ans {
				ans = f[i]
			}
		}
	}
	fmt.Println(f)
	return
}
