package main

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"middle", []int{2, 3, 4}, 6, []int{1, 3}},
		{"negative", []int{-1, 0}, -1, []int{1, 2}},
	}
	for _, tt := range tests {
		got := TwoSumII(tt.numbers, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: TwoSumII(%v, %d) = %#v, want %#v", tt.name, tt.numbers, tt.target, got, tt.want)
		}
	}
}
