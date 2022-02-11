package main

import "fmt"

func main() {
	fmt.Println("welcome to the guess the number game")
	fmt.Println("Please guess a number (0-100):")
	var gnum int = 67

	fmt.Scan(&gnum)

	if gnum < 67 {
		fmt.Println("go higher to match the secret number")
	} else if gnum > 67 {
		fmt.Println("go lower to match the secret number")
	} else {
		fmt.Println("Yay, You are a Genius :)")
	}
}
