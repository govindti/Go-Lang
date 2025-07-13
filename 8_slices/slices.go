package main

import "fmt"

func main() {
	// slice => dynamic
	// most used contruct
	// useful methods
	// uninitilized slice is nil

	// var nums []int

	// fmt.Println(nums)

	// fmt.Println(nums == nil)

	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 10)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println((cap(nums)))

	// nums = append(nums, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// fmt.Println(nums)
	// fmt.Println((cap(nums)))

	// 	nums := []int{}
	// 	fmt.Println(nums)
	// 	fmt.Println(len(nums))
	// 	fmt.Println((cap(nums)))

	// nums = append(nums, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// nums[1] = 10
	// fmt.Println(nums)
	// fmt.Println((cap(nums)))

	//copy func
	var nums = make([]int, 0, 5)
	var nums2 = make([]int, len(nums))

	nums = append(nums, 1)
	fmt.Println(nums, nums2)

	copy(nums2, nums)

	fmt.Println(nums, nums2)
}
