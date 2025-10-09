// package main

// import (
// 	"fmt"
// )

// // func printSlice[T any](items []T) {
// // func printSlice[T interface{}](items []T) {
// // func printSlice[T comparable{}](items []T) {
// func printSlice[T bool | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	// can't pass string or any other dt
// 	// printSlice([]string{"go", "js", "ts"})

// 	printSlice([]string{"go", "js", "ts"})
// 	printSlice([]bool{true, false, true})
// }

package main

import "fmt"

// can't use string bool in that stack
type stack[T any] struct {
	elements []T
}

func main() {
	myStack := stack[string]{
		elements: []string{"go", "ts", "js"},
	}
	fmt.Println(myStack)
}
