package main

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"missing", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single missing", []int{1}, 0, -1},
	}
	for _, tt := range tests {
		got := SearchRotated(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("%s: SearchRotated(%v, %d) = %d, want %d", tt.name, tt.nums, tt.target, got, tt.want)
		}
	}
}
