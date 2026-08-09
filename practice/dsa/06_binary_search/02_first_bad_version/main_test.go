package main

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		bad  int
	}{
		{"classic", 5, 4},
		{"first", 1, 1},
		{"later", 10, 8},
	}
	for _, tt := range tests {
		got := FirstBadVersion(tt.n, func(v int) bool { return v >= tt.bad })
		if got != tt.bad {
			t.Errorf("%s: FirstBadVersion(%d) = %d, want %d", tt.name, tt.n, got, tt.bad)
		}
	}
}
