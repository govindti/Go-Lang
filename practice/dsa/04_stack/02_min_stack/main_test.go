package main

import "testing"

func TestMinStack(t *testing.T) {
	t.Run("sequence", func(t *testing.T) {
		ms := NewMinStack()
		ms.Push(-2)
		ms.Push(0)
		ms.Push(-3)
		if got := ms.GetMin(); got != -3 {
			t.Errorf("GetMin() after pushes = %d, want -3", got)
		}
		ms.Pop()
		if got := ms.Top(); got != 0 {
			t.Errorf("Top() after pop = %d, want 0", got)
		}
		if got := ms.GetMin(); got != -2 {
			t.Errorf("GetMin() after pop = %d, want -2", got)
		}
	})

	t.Run("single", func(t *testing.T) {
		ms := NewMinStack()
		ms.Push(5)
		if got := ms.GetMin(); got != 5 {
			t.Errorf("GetMin() = %d, want 5", got)
		}
		if got := ms.Top(); got != 5 {
			t.Errorf("Top() = %d, want 5", got)
		}
		ms.Pop()
	})
}
