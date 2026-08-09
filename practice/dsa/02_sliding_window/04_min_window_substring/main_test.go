package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"classic", "ADOBECODEBANC", "ABC", "BANC"},
		{"single", "a", "a", "a"},
		{"impossible", "a", "aa", ""},
		{"bigger t", "ab", "abc", ""},
	}
	for _, tt := range tests {
		got := MinWindow(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("%s: MinWindow(%q, %q) = %q, want %q", tt.name, tt.s, tt.t, got, tt.want)
		}
	}
}
