package main

import "fmt"

func main() {
	// in go there are only loop named `for` but we can implement while loop using `for`

	// while
	// i := 1
	// for i <= 50 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop
	// for j := 0; j <= 20; j++ {
	// 	// break

	// 	if j==2{
	// 		continue
	// 	}
	// 	fmt.Println(j)
	// }

	// in range
	for i:= range 50 {
		fmt.Println(i)
	}



}
