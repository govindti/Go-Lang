package main

import (
	"reflect"
	"testing"
)

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

func listToSlice(head *ListNode) []int {
	out := make([]int, 0)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"classic", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"two", []int{1, 2}, []int{2, 1}},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range tests {
		got := listToSlice(ReverseList(buildList(tt.vals)))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: ReverseList(%v) = %v, want %v", tt.name, tt.vals, got, tt.want)
		}
	}
}
