package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"classic", "A man, a plan, a canal: Panama", true},
		{"not palindrome", "race a car", false},
		{"empty", "", true},
		{"only spaces", "   ", true},
	}
	for _, tt := range tests {
		got := IsPalindrome(tt.s)
		if got != tt.want {
			t.Errorf("%s: IsPalindrome(%q) = %v, want %v", tt.name, tt.s, got, tt.want)
		}
	}
}
