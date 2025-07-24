package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veggie list is: ", vegList)
	fmt.Println("Veggie list is: ", len(vegList))
	
}