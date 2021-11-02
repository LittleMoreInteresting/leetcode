package main

import (
	"fmt"
	"sync"
)

type MinStack struct {
	st     []int // slices
	min_st []int // min stack slices
	min    int
	top    int // slices top index
	l      sync.Mutex
}

func Constructor() MinStack {
	return MinStack{top: -1}
}

func (this *MinStack) Push(val int) {
	this.l.Lock()
	defer this.l.Unlock()
	this.st = append(this.st, val)
	this.top++
	if this.top == 0 || this.min > val {
		this.min = val
	}
	this.min_st = append(this.min_st, this.min)
}

func (this *MinStack) Pop() {
	this.l.Lock()
	defer this.l.Unlock()
	this.st = this.st[:this.top]
	this.min = this.min_st[this.top]
	this.min_st = this.min_st[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.st[this.top]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) getMinIndex() int {
	index := -1
	for i, v := range this.st {
		if index == -1 || this.st[index] > v {
			index = i
		}
	}
	return index
}

func (this *MinStack) Show() {
	fmt.Printf("st:%v\n", this.st)
	fmt.Printf("min_st:%v\n", this.min_st)
	fmt.Printf("min:%v\n", this.min)
	fmt.Printf("top:%v\n", this.top)
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Show()
	obj.Push(8)
	obj.Show()
	obj.Push(2)
	obj.Show()
	obj.Push(6)
	obj.Show()
	obj.Pop()
	obj.Show()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Printf("V:%v\n", param_3)
	fmt.Printf("V:%v\n", param_4)
}
