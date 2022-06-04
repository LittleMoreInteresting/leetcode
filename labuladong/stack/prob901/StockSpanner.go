package main

import "fmt"

type StockSpanner struct {
	st []Element
}

type Element struct {
	Price int
	Count int
}

func Constructor() StockSpanner {
	return StockSpanner{st: []Element{}}
}

func (s *StockSpanner) Next(price int) int {
	count := 1
	
	for len(s.st) > 0 && price >= s.st[0].Price {
		prev := s.st[0]
		s.Pop()
		count += prev.Count
	}
	s.Push(Element{Price: price, Count: count})

	return count
}

func (s *StockSpanner) Pop() {
	s.st = s.st[1:]
}

func (s *StockSpanner) Push(el Element) {
	temp := []Element{el}
	temp = append(temp, s.st...)
	s.st = temp
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}