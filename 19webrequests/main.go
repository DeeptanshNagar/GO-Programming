/**
 * -------------------------------------------------------------------
 * Program Title : Getting HTML code of UPES-ACM or any other url domain
 * Author        : Deeptansh Nagar
 * Language      : Go
 * Date          : 26th July 2025
 *
 * Description   :
 *   What this does now feels like a story –
 *   You send a robot to a website.
 *   The robot knocks on the door and says: “Give me your front page.”
 *   It brings back a big box of stuff (the website’s HTML code).
 *   You open the box and print everything it found.
 *
 * Key Features  :
 *   - This code teaches your program to talk directly to websites 
 *     the way a browser does, but without opening one.
 *   - It sends a request to a site like :- https://upesacm.org/,
 *     grabs the raw HTML code (the blueprint of the page), 
 *     and shows it in your terminal.
 *   - This is the first step to web scraping — letting you collect
 *     news, prices, weather, or any data automatically.
 *   - It also helps you understand how the web really works,
 *     because you see the behind-the-scenes “conversation” between 
 *     your code and the website.
 *   - In short, it’s a way to automate browsing, extract information, 
 *     and build smarter tools that do the boring work for you.
 * -------------------------------------------------------------------
 */


package main

import (
	"fmt"      // to print things on the screen
	"io/ioutil" // to read all the stuff we get from the website
	"net/http" // to talk to websites on the internet
)

// 📍 This is the website address we want to visit
const url = "https://upesacm.org/"

func main() {
	fmt.Println("📡 Sending a robot to visit the website...")

	// 🤖 Tell our robot to go knock on the website’s door
	response, err := http.Get(url)

	// 😱 If something went wrong (like the internet is down), stop everything!
	if err != nil {
		panic(err)
	}

	// 🚪 When we’re done, close the door behind us (important to be polite)
	defer response.Body.Close()

	// 📦 The website gives us a big box of code (HTML). Let’s open it and read everything.
	databytes, err := ioutil.ReadAll(response.Body)

	// 😱 Again, if opening the box fails, stop everything!
	if err != nil {
		panic(err)
	}

	// 📝 Turn the box of code (bytes) into words (a string we can read)
	content := string(databytes)

	// 🎉 Show everything we found inside the website
	fmt.Println("✅ Here’s what the robot brought back:")
	fmt.Println(content)
}



































// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )
// const url = "https://upesacm.org/"


// func main() {
// 	fmt.Println("LCO Web Request")

// 	response, err := http.Get(url)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Response is of type: %T\n", response)

// 	defer response.Body.Close() // caller's responsibility to close the connection

// 	// majority of readings are done by ioutil but not all remember it
// 	databytes, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	content := string(databytes)
// 	fmt.Println(content)
// }