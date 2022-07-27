package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	h := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	split(h)
}
func split(head *Node) {
	first := &Node{}
	h1 := first
	second := &Node{}
	h2 := second
	n := head
	sw := true
	show(head)
	for n != nil {
		if sw {
			first.Next = n
			first = first.Next
		} else {
			second.Next = n
			second = second.Next
		}
		n = n.Next
		sw = !sw
	}
	if sw {
		first.Next = nil
	} else {
		second.Next = nil
	}
	show(h1.Next)
	show(h2.Next)
}

func show(head *Node) {
	n := head
	for n != nil {
		fmt.Printf("%d ", n.Value)
		n = n.Next
	}
	fmt.Println()
}
