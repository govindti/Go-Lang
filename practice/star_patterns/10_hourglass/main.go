package main

import "fmt"

func main() {
	for _, line := range Hourglass(5) {
		fmt.Println(line)
	}
}

// Hourglass returns an hourglass shape. The top half has n rows
// and the bottom half has n-1 rows (total height 2n-1).
//
// n = 5:
//
//	*********
//	 *******
//	  *****
//	   ***
//	    *
//	   ***
//	  *****
//	 *******
//	*********
func Hourglass(n int) []string {
	return nil
}
