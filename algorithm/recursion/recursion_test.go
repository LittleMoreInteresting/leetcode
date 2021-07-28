package recursion

import (
	"testing"
)

func TestRecursion(t *testing.T) {
	r := Recursion(5)
	if r != 120 {
		t.Fatalf("期望 120; 实际：%v", r)
	}
	t.Logf("Success:%v\n", r)
}

func TestRecursionTail(t *testing.T) {
	r := RecursionTail(5, 1)
	if r != 120 {
		t.Fatalf("期望 120; 实际：%v", r)
	}
	t.Logf("Success:%v\n", r)
}

func TestFiboTail(t *testing.T) {
	r := FiboTail(5, 1, 1)
	t.Logf("FiboTail(5,1,1) result:%v", r)
}

func TestBinarySearch(t *testing.T) {
	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	t.Logf("arr:%v;find %v;result:%v", array, target, result)

	target2 := 123
	result2 := BinarySearch(array, target2, 0, len(array)-1)
	t.Logf("arr:%v;find %v;result:%v", array, target2, result2)
}
