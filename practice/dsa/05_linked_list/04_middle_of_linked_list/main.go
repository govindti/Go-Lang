package main

// ListNode ek singly linked list ka node hai.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode (LeetCode 876 - Middle of the Linked List)
// List ka middle node return karo. Do middle nodes ho to second wala
// (e.g. even length me right side ka node).
//
// Example:
//
//	[1, 2, 3, 4, 5]      -> node 3
//	[1, 2, 3, 4, 5, 6]   -> node 4
func MiddleNode(head *ListNode) *ListNode {
	return nil
}
