package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"anagram", "anagram", "nagaram", true},
		{"not anagram", "rat", "car", false},
		{"same letters diff count", "aacc", "ccac", false},
	}
	for _, tt := range tests {
		got := IsAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("%s: IsAnagram(%q, %q) = %v, want %v", tt.name, tt.s, tt.t, got, tt.want)
		}
	}
}
