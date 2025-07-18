package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In ChangeNum", num)
// }

// by ref
func changeNum(num *int) {
	*num = 5
	fmt.Println("In ChangeNum", *num)
}
func main() {
	num := 1
	changeNum(&num)
	fmt.Println("Memory Add:", &num)
	fmt.Println("After Change Num in main func", num)
}
