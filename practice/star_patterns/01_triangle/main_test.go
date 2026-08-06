package main

import (
	"reflect"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"*", "**", "***"}},
		{5, []string{"*", "**", "***", "****", "*****"}},
	}
	for _, tt := range tests {
		got := Triangle(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Triangle(%d) = %#v, want %#v", tt.n, got, tt.want)
		}
	}
}
