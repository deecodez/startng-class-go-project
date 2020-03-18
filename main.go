package main

// A CLI Guessing Game

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("The time is", time.Now())
	fmt.Println("Are You Good at Guessing, Guess the number!")

	//To generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100)// to generate number btw 1 and 100


	var guess int
	for{
		fmt.Println("Please enter your guess: ")
	fmt.Scan(&guess)
	if guess > secretNumber{
		fmt.Println("Almost there, figure too big Try Again!")

	}else if guess < secretNumber{
		fmt.Println("Figure to small, Try Again!")
	}else{
		fmt.Println("Congratulation, you guessed right!")
	}
}
}