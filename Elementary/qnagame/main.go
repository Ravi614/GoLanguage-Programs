package main

import "fmt"

func main() {

	fmt.Println("Welcome to the game!")
	fmt.Printf("Enter you name :")

	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, Enter Your Age :", name)

	// age limit to play this game
	var Age uint
	fmt.Scan(&Age)

	if Age >= 15 {
		fmt.Println("You can play the game")
		// continue
	} else {
		fmt.Println("You can't")
		return
	}

	scores := 0 // Count number of Questions attempt Score
	num_questions := 2 //  Total number of questions in this file
	
	// Display #1 Question to the Users

	fmt.Printf("Which one is more powerful RTX 3080 or RTX 3090 ?")

	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println("Your answer is: ", answer+" "+answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Yay!, Correct answer")
		scores++
	} else {
		fmt.Println("You lost this attempt!")
	}
	
	// Display #2 Question to the Users

	fmt.Printf("How many cards does a deck have ?")
	var cores uint // uint restricts users to enter -ve digits
	fmt.Scan(&cores)

	fmt.Println("You entered answer is :", cores)

	if cores == 52 {
		fmt.Println("Yohoo! You got the right answer!")
		scores++
	} else {
		fmt.Println("Bad luck!")
	}
	fmt.Printf("You scored %v out of %v \n", scores, num_questions)
	percentage := (scores / num_questions) * 100
	fmt.Printf("You scored %v%% out of 100", percentage)
}
