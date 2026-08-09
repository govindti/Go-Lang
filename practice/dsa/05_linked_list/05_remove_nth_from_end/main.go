package main

// ListNode ek singly linked list ka node hai.
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd (LeetCode 19 - Remove Nth Node From End of List)
// List ke end se nth node hatao aur new head return karo.
//
// Example:
//
//	RemoveNthFromEnd([1, 2, 3, 4, 5], 2) -> [1, 2, 3, 5]
//	RemoveNthFromEnd([1], 1)             -> []
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return nil
}
