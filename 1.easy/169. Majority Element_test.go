package easy

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: struct{ nums []int }{nums: []int{3, 2, 3}}, want: 3},
		{name: "Test 2", args: struct{ nums []int }{nums: []int{2, 2, 1, 1, 1, 2, 2}}, want: 2},
		{name: "Test 3", args: struct{ nums []int }{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
