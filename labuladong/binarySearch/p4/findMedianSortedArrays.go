package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	stop := totalLen / 2
	i, j := 0, 0
	rightVal, leftVal := -1, -1
	for k := 0; k <= stop; k++ {
		leftVal = rightVal
		if i < len(nums1) && (j >= len(nums2) || nums1[i] < nums2[j]) {
			rightVal = nums1[i]
			i++
		} else {
			rightVal = nums2[j]
			j++
		}
	}
	if totalLen&1 == 0 {
		return float64(rightVal+leftVal) / 2
	} else {
		return float64(rightVal)
	}

}

/**

public double findMedianSortedArrays(int[] A, int[] B) {
    int m = A.length;
    int n = B.length;
    int len = m + n;
    int left = -1, right = -1;
    int aStart = 0, bStart = 0;
    for (int i = 0; i <= len / 2; i++) {
        left = right;
        if (aStart < m && (bStart >= n || A[aStart] < B[bStart])) {
            right = A[aStart++];
        } else {
            right = B[bStart++];
        }
    }
    if ((len & 1) == 0)
        return (left + right) / 2.0;
    else
        return right;
}
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{4}
	// 1 1 2 3 3 4 4 5 6 7 8 9 9
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)
}
