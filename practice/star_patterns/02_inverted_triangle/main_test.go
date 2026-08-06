package main

import (
	"reflect"
	"testing"
)

func TestInvertedTriangle(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"***", "**", "*"}},
		{5, []string{"*****", "****", "***", "**", "*"}},
	}
	for _, tt := range tests {
		got := InvertedTriangle(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("InvertedTriangle(%d) = %#v, want %#v", tt.n, got, tt.want)
		}
	}
}
