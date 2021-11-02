package main

import (
	"sync"
)

type MinStack01 struct {
	st  []int // slices
	min int   // min index
	top int   // slices top index
	l   sync.Mutex
}

func Constructor01() MinStack01 {
	return MinStack01{st: []int{}, min: -1, top: -1}
}

func (this *MinStack01) Push(val int) {
	this.l.Lock()
	defer this.l.Unlock()
	this.st = append(this.st, val)
	this.top++
	if this.min == -1 || this.st[this.min] > val {
		this.min = this.top
	}
}

func (this *MinStack01) Pop() {
	this.l.Lock()
	defer this.l.Unlock()
	this.st = this.st[:this.top]
	if this.top == this.min {
		this.min = this.getMinIndex()
	}
	this.top--
}

func (this *MinStack01) Top() int {
	return this.st[this.top]
}

func (this *MinStack01) GetMin() int {
	return this.st[this.min]
}

func (this *MinStack01) getMinIndex() int {
	index := -1
	for i, v := range this.st {
		if index == -1 || this.st[index] > v {
			index = i
		}
	}
	return index
}
