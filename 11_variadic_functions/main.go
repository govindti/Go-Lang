package main

import "fmt"

// interface{} means any type
func main() {
	result := sum(1, 2, 2000)
	fmt.Println(result)

}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}
