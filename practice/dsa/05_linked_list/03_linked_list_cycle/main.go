package main

// ListNode ek singly linked list ka node hai.
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle (LeetCode 141 - Linked List Cycle)
// Check karo ki list me cycle hai ya nahi. Cycle means koi node ke Next
// pehle ke kisi node ko point karta hai.
//
// Example:
//
//	3 -> 2 -> 0 -> -4, jahan -4 wapas index 1 pe jata hai  -> true
//	1 -> 2 (2 wapas index 0 pe)                             -> true
//	1 -> 2 -> 3 (koi cycle nahi)                            -> false
func HasCycle(head *ListNode) bool {
	return false
}
