package main

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"classic", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"single", []int{1}, 1, []int{1}},
		{"tie", []int{1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for _, tt := range tests {
		got := TopKFrequent(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: TopKFrequent(%v, %d) = %#v, want %#v", tt.name, tt.nums, tt.k, got, tt.want)
		}
	}
}
