package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	// defer queue execution - World One Two (So, last in first out/print (LIFO)) 
	// Hello -> Two -> One -> World

	myDefer()
}
func myDefer() { 
	for i := 0; i < 5; i++ { 
		defer fmt.Println(i) // Basically, Good trick to print reverse numbers
		// Note :- adding defer doesnt allow that particular line print immediately, it stores the data into a stack than LIFO comes into the scene.
	}
}

// output guess -
// Hello 
// 0
// 1
// 2
// 3
// 4
// Two
// One
// World

// output guess after adding defer inside loop of myDefer() func
// Hello
// 4
// 3
// 2
// 1
// 0
// Two
// One
// World