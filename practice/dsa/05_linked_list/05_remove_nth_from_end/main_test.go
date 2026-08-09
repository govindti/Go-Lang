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

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
		want []int
	}{
		{"classic", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"remove only", []int{1}, 1, []int{}},
		{"remove last", []int{1, 2}, 1, []int{1}},
	}
	for _, tt := range tests {
		got := listToSlice(RemoveNthFromEnd(buildList(tt.vals), tt.n))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: RemoveNthFromEnd(%v, %d) = %v, want %v", tt.name, tt.vals, tt.n, got, tt.want)
		}
	}
}
