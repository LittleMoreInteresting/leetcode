package stackqueue

import "sync"

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	NextNode *LinkNode
	Value    string
}

func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
		return
	}

	perNode := stack.root

	newNode := new(LinkNode)
	newNode.Value = v
	newNode.NextNode = perNode
	stack.root = newNode
	stack.size = stack.size + 1
}
