package main

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"multiply", []string{"2", "1", "+", "3", "*"}, 9},
		{"division", []string{"4", "13", "5", "/", "+"}, 6},
		{"long", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tt := range tests {
		got := EvalRPN(tt.tokens)
		if got != tt.want {
			t.Errorf("%s: EvalRPN(%v) = %d, want %d", tt.name, tt.tokens, got, tt.want)
		}
	}
}
