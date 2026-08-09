package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"missing", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"single found", []int{5}, 5, 0},
		{"single missing", []int{5}, 3, -1},
	}
	for _, tt := range tests {
		got := Search(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("%s: Search(%v, %d) = %d, want %d", tt.name, tt.nums, tt.target, got, tt.want)
		}
	}
}
