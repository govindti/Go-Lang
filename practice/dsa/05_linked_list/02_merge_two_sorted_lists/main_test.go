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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{"classic", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"both empty", []int{}, []int{}, []int{}},
		{"one empty", []int{}, []int{0}, []int{0}},
	}
	for _, tt := range tests {
		got := listToSlice(MergeTwoLists(buildList(tt.list1), buildList(tt.list2)))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: MergeTwoLists(%v, %v) = %v, want %v", tt.name, tt.list1, tt.list2, got, tt.want)
		}
	}
}
