package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
	fmt.Println(maxSlidingWindow([]int{1},1))
	fmt.Println(maxSlidingWindow([]int{1,-1},1))
	fmt.Println(maxSlidingWindow([]int{9,10,9,-7,-4,-8,2,-6},5))
}

func maxSlidingWindow(nums []int, k int) []int {
	window := &MonotonicQueue{L:list.New(),}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if i < (k - 1) {
			window.Push(nums[i])
			continue
		}
		window.Push(nums[i])
		//fmt.Println(nums[i - k + 1],nums[i])
		
		max := window.Max();
		result = append(result, max)
		window.Pop(nums[i - k + 1])
	}
	return result
}

type MonotonicQueue struct {
	L *list.List
}

func (m *MonotonicQueue) Pop(n int) {
	e := m.L.Front()
	if( e.Value.(int) == n) {
		m.L.Remove(e)
	}
}

func (m *MonotonicQueue) Push(n int) {
	
	for e := m.L.Back(); e != nil && e.Value.(int) < n; e = m.L.Back() {
		m.L.Remove(e)
	}
	m.L.PushBack(n)
}

func (m *MonotonicQueue) Max() int {
	return m.L.Front().Value.(int)
}

func (m *MonotonicQueue) Show() {
	for e := m.L.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	fmt.Println()
}

func (m *MonotonicQueue) Min() int {
	return m.L.Back().Value.(int)
}

func (m *MonotonicQueue) Size() int {
	return m.L.Len()
}