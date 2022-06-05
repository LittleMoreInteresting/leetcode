package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	obj.Push(2)
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
	fmt.Println(obj.st,obj.size)
}

type MyStack struct {
	st   []int
	size int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.st = append(s.st, x)
	s.size++
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return -1
	}
	top := s.st[s.size-1]
	s.st = s.st[:s.size-1]
	s.size--
	return top
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return -1
	}
	return s.st[s.size-1]
}

func (s *MyStack) Empty() bool {
	return s.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */