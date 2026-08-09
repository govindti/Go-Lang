package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"classic", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"none", []int{0, 1, 1}, [][]int{}},
		{"empty", []int{}, [][]int{}},
	}
	for _, tt := range tests {
		got := ThreeSum(tt.nums)
		if len(tt.want) == 0 {
			if len(got) != 0 {
				t.Errorf("%s: ThreeSum(%v) = %#v, want empty result", tt.name, tt.nums, got)
			}
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: ThreeSum(%v) = %#v, want %#v", tt.name, tt.nums, got, tt.want)
		}
	}
}
