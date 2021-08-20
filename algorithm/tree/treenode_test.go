package tree

import (
	"fmt"
	"testing"
)

func TestMidOrder(test *testing.T) {
	t := MockTree()
	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
}
