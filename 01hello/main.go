package main

import "fmt"

func main() {
	fmt.Println("Hello, World!"); // Putting semi colon is not that much necessary
}

// Lexer in golang and Types

// Semicolons in Go - you typically don't need to write them explicitly. Go has a unique approach called "automatic semicolon insertion." The lexer automatically adds semicolons at the end of lines when certain conditions are met, such as when a line ends with:

// An identifier
// A literal (number, string, etc.)
// Keywords like break, continue, return
// Operators like ++, --
// Closing brackets }, ), ]

// So this code works without semicolons:
// gofunc main() {
//     x := 10
//     fmt.Println(x)
//     return
// }

// The lexer automatically treats it as:
// gofunc main() {
//     x := 10;
//     fmt.Println(x);
//     return;
// }

// You can still write semicolons explicitly if you want, and they're required when putting multiple statements on one line:
// x := 10; y := 20; fmt.Println(x, y)

// This automatic insertion is why Go's formatting rules are strict about brace placement - putting the opening brace on the next line would cause the lexer to insert an unwanted semicolon.