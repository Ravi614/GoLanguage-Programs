package main

import "fmt"

// program such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17
func main() {
	fmt.Println("Welcome to the divison zone")
	var number int
	fmt.Printf("Enter the number: ")
	fmt.Scan(&number)

	for number := 1; number <= 17; number++ {
		if number%3 == 0 || number%5 == 0 {
			fmt.Println(number)
		}
	}
}
