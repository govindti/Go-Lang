package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"not sorted", []int{3, 2, 4}, 6, []int{1, 2}},
		{"same value twice", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: TwoSum(%v, %d) = %#v, want %#v", tt.name, tt.nums, tt.target, got, tt.want)
		}
	}
}
