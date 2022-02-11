package main

// Write a program to count the total digit entered by user

import "fmt"

func main() {
	number := 0
	fmt.Print("Please enter number: ")
	fmt.Scanln(&number)
	fmt.Printf("Digits count of %d is %d", number, iterativeDigitsCount(number))
}

func iterativeDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10 // 456 = 456/10 = 45.6 | Since it's a float value hence it will consider 45 and we instructed program to return int
		count += 1
	}
	return count
}
