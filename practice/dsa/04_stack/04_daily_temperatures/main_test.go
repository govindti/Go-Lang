package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"classic", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"increasing", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"decreasing", []int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, tt := range tests {
		got := DailyTemperatures(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: DailyTemperatures(%v) = %#v, want %#v", tt.name, tt.nums, got, tt.want)
		}
	}
}
