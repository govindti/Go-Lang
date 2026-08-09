package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"classic", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"two", []int{1, 1}, 1},
		{"walls", []int{4, 3, 2, 1, 4}, 16},
	}
	for _, tt := range tests {
		got := MaxArea(tt.height)
		if got != tt.want {
			t.Errorf("%s: MaxArea(%v) = %d, want %d", tt.name, tt.height, got, tt.want)
		}
	}
}
