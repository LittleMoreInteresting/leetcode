package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	nums := strings.Split(str, " ")
	l, _ := strconv.Atoi(nums[0])
	mepList := mepList{l: NewNode(), m: make(map[string]*Node)}
	head := nums[1]
	hNode := &Node{Val: head}
	hNode.Next = mepList.l.Next
	hNode.Prve = mepList.l
	mepList.l.Next = hNode
	mepList.m[head] = hNode
	del := nums[len(nums)-1]

	for i := 0; i < l-1; i++ {
		val := nums[2+2*i]
		prve := nums[3+2*i]
		pNode := mepList.m[prve]
		hNode := &Node{Val: val, Prve: pNode}
		hNode.Next = pNode.Next
		pNode.Next.Prve = hNode
		pNode.Next = hNode

		mepList.m[val] = hNode
	}
	delNode := mepList.m[del]
	delNode.Next.Prve = delNode.Prve
	delNode.Prve.Next = delNode.Next

	for next := mepList.l.Next; next != nil; next = next.Next {
		fmt.Print(next.Val)
		if next.Next != nil {
			fmt.Print(" ")
		}
	}
}

type mepList struct {
	l *Node
	m map[string]*Node
}

type Node struct {
	Val  string
	Prve *Node
	Next *Node
}

func NewNode() *Node {
	head := &Node{}
	return &Node{Prve: head, Next: head}
}
