package main

import (
	"fmt"
)

// Write a program that asks the user for their name and greets them with their name.

func main() {

	fmt.Println("Welcome to a new UserApp")

	fmt.Printf("Enter your name :")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, Welcome to new userapp", name)
}
