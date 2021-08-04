package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "empty", args: args{nums: []int{}, target: 4}, want: -1},
		{name: "not fount", args: args{nums: []int{1, 2, 3, 4, 5, 6}, target: 9}, want: -1},
		{name: "fount", args: args{nums: []int{1, 2, 3, 4, 5, 6, 9, 23, 24}, target: 9}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
