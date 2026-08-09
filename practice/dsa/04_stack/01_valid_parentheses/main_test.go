package main

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"simple", "()", true},
		{"all types", "()[]{}", true},
		{"mismatched", "(]", false},
		{"wrong nesting", "([)]", false},
		{"unclosed", "(", false},
	}
	for _, tt := range tests {
		got := IsValidParentheses(tt.s)
		if got != tt.want {
			t.Errorf("%s: IsValidParentheses(%q) = %v, want %v", tt.name, tt.s, got, tt.want)
		}
	}
}
