package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Game of Guessing")
	fmt.Println("One number random be sorted from 1 to 100. Try to guess it. You have 10 chances.")

	scanner := bufio.NewScanner(os.Stdin)
	attemps := [10]int64{}
	x := rand.Int64N(101)
	for i := range attemps {
		fmt.Println("Enter your guess: ")

		scanner.Scan()
		attemp := scanner.Text()

		attemp = strings.TrimSpace(attemp)

		attempInteger, err := strconv.ParseInt(attemp, 10, 64)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		switch {
		case attempInteger < x:
			fmt.Println("Your guess is too low.", attempInteger)
		case attempInteger > x:
			fmt.Println("Your guess is too high.", attempInteger)
		case attempInteger == x:
			fmt.Printf("Congratulations! You guessed the number %d\n"+
				"You guessed the number in %d attempts.\n"+
				"These were your attempts: %v\n", x, i+1, attemps[:i])
			return
		}

		attemps[i] = attempInteger
	}

	fmt.Printf("You didn't guess the number. The number was %d\n", x)
	fmt.Println("Your guesses were: ", attemps)
}
