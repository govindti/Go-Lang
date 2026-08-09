package main

import "testing"

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"classic", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"single", []int{5}, 1, 5.0},
		{"negative", []int{0, 1, 1, 3, 3}, 4, 2.0},
	}
	for _, tt := range tests {
		got := FindMaxAverage(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("%s: FindMaxAverage(%v, %d) = %v, want %v", tt.name, tt.nums, tt.k, got, tt.want)
		}
	}
}
