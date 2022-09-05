package p135

func candy(ratings []int) int {
	l := len(ratings)
	if l < 2 {
		return l
	}
	total := make([]int, l)
	total[0] = 1
	for i := 1; i < l; i++ {
		total[i] = 1
		if ratings[i] > ratings[i-1] {
			total[i] = total[i-1] + 1
		}
	}
	for i := l - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			total[i-1] = Max(ratings[i-1], ratings[i]+1)
		}
	}
	return Sum(total)
}

func Sum(num []int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/candy
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func candy1(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
