package easy

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"Example 1", args{"III"}, 3},
		{"Example 2", args{"IV"}, 4},
		{"Example 3", args{"IX"}, 9},
		{"Example 4", args{"LVIII"}, 58},
		{"Example 5", args{"MCMXCIV"}, 1994},
		{"Example 6", args{"DCXXI"}, 621},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
