package main

import "fmt"

func main() {
	for _, line := range HollowTriangle(5) {
		fmt.Println(line)
	}
}

// HollowTriangle returns a right triangle that is hollow inside.
// The first two rows are always solid, the last row is always solid.
//
// n = 5:
//
//	*
//	**
//	* *
//	*  *
//	*****
func HollowTriangle(n int) []string {
	return nil
}
