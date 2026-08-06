package main

import "fmt"

func main() {
	for _, line := range Diamond(5) {
		fmt.Println(line)
	}
}

// Diamond returns a diamond shape. The top half has n rows
// and the bottom half has n-1 rows (total height 2n-1).
//
// n = 5:
//
//	    *
//	   ***
//	  *****
//	 *******
//	*********
//	 *******
//	  *****
//	   ***
//	    *
func Diamond(n int) []string {
	return nil
}
