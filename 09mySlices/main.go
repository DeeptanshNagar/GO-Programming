package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList) // [Apple Tomato Peach Mango Banana]

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList) // [Tomato Peach Mango Banana]

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList) // [Tomato Mango]

	// highScores := make([]int, 4)
	// highScores[0] = 200
	// highScores[1] = 900
	// highScores[2] = 400
	// highScores[3] = 800
	// highScores[4] = 700
	// highScores = append(highScores, 555, 666, 234)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))


	// vid 15 -
	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript",  "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // append method reallocates the memory
	fmt.Println(courses)
} 