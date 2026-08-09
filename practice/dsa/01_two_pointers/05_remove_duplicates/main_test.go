package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"basic", []int{1, 1, 2}, 2},
		{"long", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{"empty", []int{}, 0},
		{"all same", []int{7, 7, 7, 7}, 1},
	}
	for _, tt := range tests {
		got := RemoveDuplicates(tt.nums)
		if got != tt.want {
			t.Errorf("%s: RemoveDuplicates(%v) = %d, want %d", tt.name, tt.nums, got, tt.want)
		}
	}
}
