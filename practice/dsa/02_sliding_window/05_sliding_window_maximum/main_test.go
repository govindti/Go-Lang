package main

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"classic", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"single", []int{1}, 1, []int{1}},
		{"negative", []int{1, -1}, 1, []int{1, -1}},
	}
	for _, tt := range tests {
		got := MaxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: MaxSlidingWindow(%v, %d) = %#v, want %#v", tt.name, tt.nums, tt.k, got, tt.want)
		}
	}
}
