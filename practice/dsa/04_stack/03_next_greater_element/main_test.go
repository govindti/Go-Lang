package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"classic", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"sorted", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{"all to last", []int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}, []int{7, 7, 7, 7, 7}},
	}
	for _, tt := range tests {
		got := NextGreaterElement(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: NextGreaterElement(%v, %v) = %#v, want %#v", tt.name, tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
