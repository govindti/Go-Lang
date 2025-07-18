package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id     string
	amount float32
	status string
	customer
	createdAt time.Time // nano second precision

}

// receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func newOder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

func main() {
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 456.4,
	// 	status: "received",
	// }

	// myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.status)

	// fmt.Println(myOrder)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println(myOrder2)

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 456.4,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	// myOrder := newOder("3", 30.2, "ho gya")
	// fmt.Println(myOrder)

	// lang := struct {
	// 	name   string
	// 	isGood bool
	// }{"GoLang", true}
	// fmt.Println(lang)

	newOrder := order{
		id:     "1",
		amount: 30,
		status: "paid",
		customer: customer{
			name:  "John",
			phone: "123456789",
		},
	}

	fmt.Println(newOrder)
}
