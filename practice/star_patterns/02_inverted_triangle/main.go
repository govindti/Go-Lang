package main

import "fmt"

func main() {
	for _, line := range InvertedTriangle(5) {
		fmt.Println(line)
	}
}

// InvertedTriangle returns the inverted (bottom-left corner) triangle of size n.
//
// n = 5:
//
//	*****
//	****
//	***
//	**
//	*
func InvertedTriangle(n int) []string {
	return nil
}
