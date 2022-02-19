package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1

	fmt.Println("This app has fixed random number between 1 to 100")
	fmt.Println("Can you guess it ?")
	// fmt.Println(target)

	// Get the input from keyboard
	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have ", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		
		// Readstring returns strings, hence convert to int either by parseFloat or Atoi

		input, err := reader.ReadString('\n')

		errorGen(err)

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		errorGen(err)

		if guess < target {
			fmt.Println("Oops, Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops, Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You Guessed it right!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number correctly. It was", target)
	}
}

// Avoid repetation of this below sode snippet

func errorGen(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
