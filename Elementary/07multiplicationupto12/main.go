package main

import "fmt"

//  Write a program that prints a multiplication table for numbers up to 12.
func main() {
	var i, j int
	fmt.Println("\nMultiplication Table of 1 to 12 are as follows: ")

	for i = 1; i <= 12; i++ {
		for j = 1; j <= 12; j++ {
			fmt.Println(i * j)
		}
		fmt.Printf("--------- \n")
	}
}
