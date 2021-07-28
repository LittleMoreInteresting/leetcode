package linkedlists

import "fmt"

type linkNode struct {
	Data     int64
	NextNode *linkNode
}

func (l *linkNode) show() {
	current := l
	space := "->"
	for {
		fmt.Printf("%v%s", current.Data, space)
		if current.NextNode != nil {
			current = current.NextNode
			continue
		}
		break
	}
}

func NewLink(data int64) *linkNode {
	return &linkNode{
		Data:     data,
		NextNode: nil,
	}
}
