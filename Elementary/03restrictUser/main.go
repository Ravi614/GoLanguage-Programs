package main

import (
	"fmt"
)

// Modify the greetuser program such that only the users Alice and Bob are greeted with their names.

func main() {

	fmt.Println("Welcome to a new UserApp")

	fmt.Printf("Enter your name :")

	var name string
	fmt.Scan(&name)

	if name == "Alice" || name == "alice" || name == "Bob" || name == "bob" {
		fmt.Printf("Hello, %v, Welcome to the app! ", name)
	} else {
		fmt.Printf("Try registering with us!")
	}
}
