package main

import "fmt"

// Write a program that asks the user for a number n and gives them the possibility to choose between computing the sum and computing the product of 1,…,n

func main() {
	fmt.Println("Welcome to the computing options app")
	var digit int
	fmt.Printf("Enter any digit of your choice: ")
	fmt.Scan(&digit)
	fmt.Println("You have entered the digit: ", digit)

	fmt.Printf("What would you like to compute? The sum (1) or the product (2): ")

	var choice string
	var count int = 1

	fmt.Scan(&choice)

	if choice == "1" {
		fmt.Println(digit + count)
		count++
	} else if choice == "2" {
		fmt.Println(digit * count)
		count++
	} else {
		return
	}
}
