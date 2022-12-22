package p1745

func countBalls(lowLimit, highLimit int) (ans int) {
	count := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := numSum(i)
		count[sum]++
		ans = max(ans, count[sum])
	}
	return
}

func numSum(n int) (sum int) {
	for x := n; x > 0; x /= 10 {
		sum += x % 10
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/solutions/1984796/he-zi-zhong-xiao-qiu-de-zui-da-shu-liang-9sfh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
