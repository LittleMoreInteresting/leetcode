package rwmap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := NewRWMap[string, int](10)
	m.Set("key1", 10)
	v1, _ := m.Get("key1")
	fmt.Println(v1)

}

func TestFnv(t *testing.T) {
	h1 := fnv32("abc")
	h2 := fnv32("abc")
	fmt.Println(h1 & 63)
	fmt.Println(h2 & 63)
}
