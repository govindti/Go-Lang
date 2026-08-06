package main

import "fmt"

func main() {
	for _, line := range Pyramid(5) {
		fmt.Println(line)
	}
}

// Pyramid returns a centered full pyramid of n rows.
//
// n = 5:
//
//	    *
//	   ***
//	  *****
//	 *******
//	*********
func Pyramid(n int) []string {
	return nil
}
