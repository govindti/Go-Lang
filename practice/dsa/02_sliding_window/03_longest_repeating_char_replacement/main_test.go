package main

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"abab", "ABAB", 2, 4},
		{"classic", "AABABBA", 1, 4},
		{"no change needed", "AAAA", 0, 4},
	}
	for _, tt := range tests {
		got := CharacterReplacement(tt.s, tt.k)
		if got != tt.want {
			t.Errorf("%s: CharacterReplacement(%q, %d) = %d, want %d", tt.name, tt.s, tt.k, got, tt.want)
		}
	}
}
