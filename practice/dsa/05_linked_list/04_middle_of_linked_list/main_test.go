package main

import "testing"

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"odd", []int{1, 2, 3, 4, 5}, 3},
		{"even", []int{1, 2, 3, 4, 5, 6}, 4},
		{"single", []int{1}, 1},
	}
	for _, tt := range tests {
		got := MiddleNode(buildList(tt.vals))
		if got == nil || got.Val != tt.want {
			val := -1
			if got != nil {
				val = got.Val
			}
			t.Errorf("%s: MiddleNode(%v) = %d, want %d", tt.name, tt.vals, val, tt.want)
		}
	}
}
