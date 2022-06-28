package main

import "fmt"

func main() {
	fs := Constructor()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
	fmt.Println(fs.Pop())
}

type FreqStack struct {
	maxFreq     int
	valueToFreq map[int]int
	freqToVal   map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		valueToFreq: make(map[int]int),
		freqToVal:   make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	f := this.valueToFreq[val]
	freq := f + 1
	temp := this.freqToVal[freq]
	if temp == nil {
		temp = make([]int, 0)
	}
	temp = append(temp, val)
	this.freqToVal[freq] = temp
	this.valueToFreq[val] = freq
	if freq > this.maxFreq {
		this.maxFreq = freq
	}
}

func (this *FreqStack) Pop() int {
	bucket := this.freqToVal[this.maxFreq]
	last := bucket[len(bucket)-1]
	bucket = bucket[:len(bucket)-1]
	this.freqToVal[this.maxFreq] = bucket
	if len(bucket) == 0 {
		this.maxFreq--
	}
	this.valueToFreq[last] -= 1
	return last
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
