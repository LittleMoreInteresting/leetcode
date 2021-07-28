package linkedlists

import "testing"

func TestNewLink(t *testing.T) {
	link := NewLink(1)
	link2 := NewLink(2)
	link3 := NewLink(3)
	link.NextNode = link2
	link2.NextNode = link3
	link.show()
}
