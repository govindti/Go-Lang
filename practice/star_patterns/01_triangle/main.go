package main

import "fmt"

func main() {
	for _, line := range Triangle(5) {
		fmt.Println(line)
	}
}

// Triangle returns the right-aligned (top-left corner) triangle of size n.
//
// n = 5:
//
//	*
//	**
//	***
//	****
//	*****
func Triangle(n int) []string {
	return nil
}
