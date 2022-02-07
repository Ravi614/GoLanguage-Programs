package main

import "fmt"

// Write a program that asks the user for a number n and prints the table of that number

func main() {
	fmt.Printf("Enter a number : ")
	var num int
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num * i)
	}
}
