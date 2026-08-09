package main

// MinStack (LeetCode 155 - Min Stack)
// Ek stack jo push, pop, top aur minimum element — sab O(1) me — support
// kare. Struct me fields khud add karo (e.g. stack + min stack).
//
// Use:
//
//	ms := NewMinStack()
//	ms.Push(-2); ms.Push(0); ms.Push(-3)
//	ms.GetMin() -> -3
//	ms.Pop()
//	ms.Top()    -> 0
//	ms.GetMin() -> -2
type MinStack struct{}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (ms *MinStack) Push(val int) {}

func (ms *MinStack) Pop() {}

func (ms *MinStack) Top() int {
	return 0
}

func (ms *MinStack) GetMin() int {
	return 0
}
