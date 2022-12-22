package p878

const mod int = 1e9 + 7

func nthMagicalNumber(n, a, b int) int {
	l := min(a, b)
	r := n * min(a, b)
	c := a / gcd(a, b) * b
	for l <= r {
		mid := (l + r) / 2
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/nth-magical-number/solutions/1983699/di-n-ge-shen-qi-shu-zi-by-leetcode-solut-6vyy/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
