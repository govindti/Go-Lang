package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"classic", "abcabcbb", 3},
		{"all same", "bbbbb", 1},
		{"subsequence", "pwwkew", 3},
		{"empty", "", 0},
	}
	for _, tt := range tests {
		got := LengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("%s: LengthOfLongestSubstring(%q) = %d, want %d", tt.name, tt.s, got, tt.want)
		}
	}
}
