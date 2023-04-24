package main

import (
	"fmt"
	"strconv"
)

func showMenu(choices []string) {
	fmt.Println("Enter the number corresponding to your choice")

	for i := 0; i < len(choices); i++ {
		fmt.Println(i+1, "->", choices[i])
	}

	var userInput string

	fmt.Scanln(&userInput)
	userChoice, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Invalid input. Please retry!")
		showMenu(choices)

		return
	}

	if userChoice < 1 || userChoice > len(choices) {
		fmt.Println("Choice unavailable. Please retry!")
		showMenu(choices)

		return
	}

	fmt.Println(userChoice)
	fmt.Println("You chose to", choices[userChoice-1])
}

func main() {
	fmt.Println("Welcome to the monty hall game!")
	showMenu([]string{"Play", "Quit"})
}
