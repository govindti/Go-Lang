package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found", []int{1, 3, 5, 6}, 5, 2},
		{"insert middle", []int{1, 3, 5, 6}, 2, 1},
		{"insert end", []int{1, 3, 5, 6}, 7, 4},
		{"insert start", []int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		got := SearchInsert(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("%s: SearchInsert(%v, %d) = %d, want %d", tt.name, tt.nums, tt.target, got, tt.want)
		}
	}
}
