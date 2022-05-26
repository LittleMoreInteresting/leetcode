package main

import "fmt"

type NumArray struct {
	//Nums    []int
	prevSum []int
}

func Constructor(nums []int) NumArray {
	prevSum := make([]int, len(nums)+1)
	for i := 1; i < len(prevSum); i++ {
		prevSum[i] = prevSum[i-1] + nums[i-1]
	}

	return NumArray{/*Nums: nums,*/ prevSum: prevSum}

}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prevSum[right+1] - this.prevSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	// 0 -2 -2 1 -4 -2 -1
	obj := Constructor(nums)
	fmt.Println(obj.prevSum)
	param_1 := obj.SumRange(0, 2)
	param_2 := obj.SumRange(2, 5)
	param_3 := obj.SumRange(0, 5)
	fmt.Println(param_1,param_2,param_3)
}