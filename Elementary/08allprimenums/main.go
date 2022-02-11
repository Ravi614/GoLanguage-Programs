package main

import (
	"fmt"
	"math/big"
)

// Write a program that prints all prime numbers.
// (Note: if your programming language does not support arbitrary size numbers, printing all primes up to the largest number you can easily represent is fine too.)

func main() {
	for i := 2; i < 100; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			fmt.Printf("%d is prime nos \n", i)
		}
	}
}

// TODO: Without using package
