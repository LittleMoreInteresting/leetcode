package main

import "fmt"

func main() {
	fmt.Println(numTilings(3))
	fmt.Println(numTilings1(3))
}
func numTilings(n int) int {
	const mod int = 1e9 + 7
	dp_0, dp_1, dp_2, dp_3 := 0, 0, 0, 1
	for i := 1; i <= n; i++ {
		temp_0, temp_1, temp_2, temp_3 := dp_0, dp_1, dp_2, dp_3
		dp_0 = temp_3
		dp_1 = (temp_0 + temp_2) % mod
		dp_2 = (temp_0 + temp_1) % mod
		dp_3 = (((temp_0+temp_1)%mod+temp_2)%mod + temp_3) % mod
	}
	return dp_3
}

func numTilings1(n int) int {
	const mod int = 1e9 + 7
	dp := make([][4]int, n+1)
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (((dp[i-1][0]+dp[i-1][1])%mod+dp[i-1][2])%mod + dp[i-1][3]) % mod
	}
	return dp[n][3]
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/domino-and-tromino-tiling/solutions/1962465/duo-mi-nuo-he-tuo-mi-nuo-ping-pu-by-leet-7n0j/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
