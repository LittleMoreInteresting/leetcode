package p1233

import (
	"fmt"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	re := removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"})
	fmt.Println(re)
}
func Test_isChild(t *testing.T) {
	re := isChild("/a/b/c", "/a/b")
	fmt.Println(re)
}
