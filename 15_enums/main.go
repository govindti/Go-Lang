package main

import "fmt"

type OrderStatus int

// iota if int
const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

//  if string
// const (
// 	Received OrderStatus = "received"
// 	Confirmed ="confirmed"
// 	Prepared ="prepaerd"
// 	Delivered ="delivered"
// )

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order Status to", status)
}

func main() {
	changeOrderStatus(Delivered)
}
