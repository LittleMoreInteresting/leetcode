package linkedlists

import (
	"fmt"
	"testing"
)

func TestNewRing(t *testing.T) {
	r := &Ring{Value: -1}

	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	r.Link(&Ring{Value: 5})

	node := r
	for {
		fmt.Println(node.Value)
		node = node.Next()
		if node == r {
			break
		}
	}

	temp := r.Unlink(3)
	fmt.Println("------")

	node = temp
	for {
		fmt.Println(node.Value)
		node = node.Next()
		if node == temp {
			break
		}
	}
}
