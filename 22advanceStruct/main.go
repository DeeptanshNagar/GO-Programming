package main

import (
	"fmt"
	// "time"
)

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// // constructor
// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here...
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status, 
// 	}

// 	return &myOrder
// }

// // receiver type 
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// // important note :-
// // if we want to modify the value then we have to necessarily use the pointer *.
// // if we want to get the value then using the pointer is not necessary.

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	// inline struct // used for that instance only
	language := struct {
		name string
		isGood bool
	} {"golang", true}
	fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received", 

	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println("Order struct: ", myOrder.getAmount()) // 0 if not set
	// fmt.Println("Order struct: ", myOrder)


	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order {
	// 	id: "2",
	// 	amount: 100,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }
	// myOrder.status = "paid"
	// fmt.Println("Order struct: ", myOrder)
	// fmt.Println("Order struct: ", myOrder2)
}