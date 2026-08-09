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

func buildCyclicList(vals []int, pos int) *ListNode {
	head := buildList(vals)
	if pos < 0 {
		return head
	}
	var tail, cycleNode *ListNode = head, head
	idx := 0
	for tail.Next != nil {
		if idx == pos {
			cycleNode = tail
		}
		tail = tail.Next
		idx++
	}
	if idx == pos {
		cycleNode = tail
	}
	tail.Next = cycleNode
	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{"cycle at 1", []int{3, 2, 0, -4}, 1, true},
		{"cycle at 0", []int{1, 2}, 0, true},
		{"no cycle", []int{1}, -1, false},
		{"empty", []int{}, -1, false},
	}
	for _, tt := range tests {
		got := HasCycle(buildCyclicList(tt.vals, tt.pos))
		if got != tt.want {
			t.Errorf("%s: HasCycle(%v, pos=%d) = %v, want %v", tt.name, tt.vals, tt.pos, got, tt.want)
		}
	}
}
