package main

import "fmt"

func main() {
	for _, line := range Square(5) {
		fmt.Println(line)
	}
}

// Square returns an n x n solid square of stars.
//
// n = 5:
//
//	*****
//	*****
//	*****
//	*****
//	*****
func Square(n int) []string {
	return nil
}
