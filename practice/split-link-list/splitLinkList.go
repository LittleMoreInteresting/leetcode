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
			h1.Next = n
			h1 = h1.Next
		} else {
			h2.Next = n
			h2 = h2.Next
		}
		temp := n.Next
		n.Next = nil
		n = temp
		sw = !sw
	}
	/*if sw {
		h1.Next = nil
	} else {
		h2.Next = nil
	}*/
	show(first.Next)
	show(second.Next)
}

func show(head *Node) {
	n := head
	for n != nil {
		fmt.Printf("%d ", n.Value)
		n = n.Next
	}
	fmt.Println()
}
