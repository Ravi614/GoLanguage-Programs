package main

import "fmt"

func main() {
	originalCount := 10 // := is called short variable declarations
	eatenCount := 4
	fmt.Println("I started with", originalCount, "apples")
	fmt.Println("Some jerk ate", eatenCount, "apples")

	fmt.Println("There are", originalCount-eatenCount, "apples left")

}
