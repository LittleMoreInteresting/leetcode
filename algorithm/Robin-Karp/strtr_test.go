package Robin_Karp

import "testing"

func TestSearchStr(t *testing.T) {
	type args struct {
		s      string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "has", args: args{"abcdefg", "cd"}, want: 2},
		{name: "has", args: args{"abcdefg", "gh"}, want: -1},
		{name: "has", args: args{"abcdefg", "a"}, want: 0},
		{name: "has", args: args{"abcdefg", "abcdefg"}, want: 0},
		{name: "has", args: args{"abcdefg", "abcdefgh"}, want: -1},
		{name: "has", args: args{"abcdefg", "a"}, want: 0},
		{name: "has", args: args{"abcdefg", "b"}, want: 1},
		{name: "has", args: args{"abcdefg", "c"}, want: 2},
		{name: "has", args: args{"abcdefg", "g"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchStr(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("SearchStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
