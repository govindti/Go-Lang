package main

import "fmt"

func main() {
	for _, line := range HollowSquare(5) {
		fmt.Println(line)
	}
}

// HollowSquare returns an n x n square that is hollow inside.
//
// n = 5:
//
//	*****
//	*   *
//	*   *
//	*   *
//	*****
func HollowSquare(n int) []string {
	return nil
}
