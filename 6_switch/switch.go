package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("1")

	// case 2:
	// 	fmt.Println("2")

	// case 3:
	// 	fmt.Println("3")

	// case 4:
	// 	fmt.Println("4")

	// case 5:
	// 	fmt.Println("5")
	// default:
	// 	fmt.Println("Others")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's WeekEnd")

	// default:
	// 	fmt.Println("It's Weekday")
	// }


	// type switch
	whomI := func (i interface{})  {
		switch i:= i.(type) {
			
		}
	}
}
