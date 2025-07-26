package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Writing the file
func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")
	
	// if err != nil {
	// 	panic(err)
	// }
	// use this as alternative -
	checkNilErr(err)

	length, err  := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

// Reading the file 
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // file isnt being read in string format, its being read in byte format.
	checkNilErr(err)

	// fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}