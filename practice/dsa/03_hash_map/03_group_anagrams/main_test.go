package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{"classic", []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"ate", "eat", "tea"}, {"nat", "tan"}}},
		{"empty string", []string{""}, [][]string{{""}}},
		{"single", []string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		got := GroupAnagrams(tt.strs)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: GroupAnagrams(%v) = %#v, want %#v", tt.name, tt.strs, got, tt.want)
		}
	}
}
