package main

import "fmt"

// zeroed values concept (int => 0, bool => false, string => "")
func main() {
	// var nums [4]int

	// nums[0] = 1

	// fmt.Println(nums)

	// var boolArray [3]bool
	// fmt.Println(boolArray)

	// var strArray [3]string
	// fmt.Println(strArray)

	//with initial values.....................

	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	// 2d array.....................
	nums := [2][2]int{{1, 3}, {2, 4}}
	fmt.Println(nums)

}
