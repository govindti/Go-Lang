package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"rotated", []int{3, 4, 5, 1, 2}, 1},
		{"more rotated", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"not rotated", []int{11, 13, 15, 17}, 11},
	}
	for _, tt := range tests {
		got := FindMin(tt.nums)
		if got != tt.want {
			t.Errorf("%s: FindMin(%v) = %d, want %d", tt.name, tt.nums, got, tt.want)
		}
	}
}
