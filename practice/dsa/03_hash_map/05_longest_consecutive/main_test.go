package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic", []int{100, 4, 200, 1, 3, 2}, 4},
		{"long", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"empty", []int{}, 0},
	}
	for _, tt := range tests {
		got := LongestConsecutive(tt.nums)
		if got != tt.want {
			t.Errorf("%s: LongestConsecutive(%v) = %d, want %d", tt.name, tt.nums, got, tt.want)
		}
	}
}
