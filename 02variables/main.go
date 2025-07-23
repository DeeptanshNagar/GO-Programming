package main

import "fmt"

func main() {
	var username string = "deeptansh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	fmt.Println()

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	fmt.Println()

	var smallValue uint8 = 255 // 256 will give you an error bcoz range is 0-255
	// uint16 -> 0 to 65535 (0 to 2¹⁶ - 1)
	// uint 32 -> 0 to 4,294,967,295 (0 to 2³² - 1)
	// unit 64 -> 0 to 18,446,744,073,709,551,615 (0 to 2⁶⁴ - 1)
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	fmt.Println()

	var smallFloat float32 = 255.4553535355536
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	fmt.Println()

	var smallFloatt float64 = 255.4553535355536 // increases the precision
	fmt.Println(smallFloatt)
	fmt.Printf("Variable is of type: %T \n", smallFloatt)

	fmt.Println()

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	fmt.Println()

	// implicit type - means if you wont declare data type its-fine!
	var website = "learncodeonline.in"
	fmt.Println(website)

	// You can't change it later to another data type except string because it was initialized with a string data type previously.

	// website = 300 // this'll throw an error
	// fmt.Println(website)

	fmt.Println()

	// no var style
	numberOfUser := 300000.0 // you can change it to any data type and still it works fine
	// := (You are only allowed to use the walrus operator (:=) inside methods/functions.)
	fmt.Println(numberOfUser)


}
