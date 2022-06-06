package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{{5,4},{6,4},{6,7},{2,3}}
	fmt.Println(maxEnvelopes(envelopes))
	envelopes = [][]int{{1,1},{1,1},{1,1}}
	fmt.Println(maxEnvelopes(envelopes))
}
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
        a, b := envelopes[i], envelopes[j]
        return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
    })

    f := []int{}
    for _, e := range envelopes {
        h := e[1]
        if i := sort.SearchInts(f, h); i < len(f) {
            f[i] = h
        } else {
            f = append(f, h)
        }
    }
	fmt.Println(f)
    return len(f)
}
func maxEnvelopes_dp(envelopes [][]int) int {
	l := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
        a, b := envelopes[i], envelopes[j]
        return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
    })
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
                dp[i] = max(dp[i], dp[j]+1)
            }
		}
	}
	return max(dp...)
}

func max(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}

func maxEnvelopes_gf(envelopes [][]int) int {
    n := len(envelopes)
    if n == 0 {
        return 0
    }

    sort.Slice(envelopes, func(i, j int) bool {
        a, b := envelopes[i], envelopes[j]
        return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
    })

    f := make([]int, n)
    for i := range f {
        f[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if envelopes[j][1] < envelopes[i][1] {
                f[i] = max(f[i], f[j]+1)
            }
        }
    }
    return max(f...)
}

