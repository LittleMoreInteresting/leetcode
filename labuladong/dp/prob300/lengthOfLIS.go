package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}

func lengthOfLIS(nums []int) int {
	//base case
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLISBS(nums []int) int {
	//TODO 二分查找
	l := len(nums)
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// int[] top = new int[nums.length];
// // 牌堆数初始化为 0
// int piles = 0;
// for (int i = 0; i < nums.length; i++) {
// 	// 要处理的扑克牌
// 	int poker = nums[i];

// 	/***** 搜索左侧边界的二分查找 *****/
// 	int left = 0, right = piles;
// 	while (left < right) {
// 		int mid = (left + right) / 2;
// 		if (top[mid] > poker) {
// 			right = mid;
// 		} else if (top[mid] < poker) {
// 			left = mid + 1;
// 		} else {
// 			right = mid;
// 		}
// 	}
// 	/*********************************/
	
// 	// 没找到合适的牌堆，新建一堆
// 	if (left == piles) piles++;
// 	// 把这张牌放到牌堆顶
// 	top[left] = poker;
// }
// // 牌堆数就是 LIS 长度
// return piles;
