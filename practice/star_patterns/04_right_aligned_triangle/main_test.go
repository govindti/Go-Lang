package main

import (
	"reflect"
	"testing"
)

func TestRightAlignedTriangle(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"*"}},
		{3, []string{"  *", " **", "***"}},
		{5, []string{"    *", "   **", "  ***", " ****", "*****"}},
	}
	for _, tt := range tests {
		got := RightAlignedTriangle(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RightAlignedTriangle(%d) = %#v, want %#v", tt.n, got, tt.want)
		}
	}
}
