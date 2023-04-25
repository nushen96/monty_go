package main

import (
	"fmt"
	"strconv"
  "math/rand"
  "errors"
)

type Stats struct {
  Played int
  Won int
}

var gameModeStats Stats = Stats{Played: 0, Won: 0}

func calculatePercentage(num1 int, num2 int) (float64, error) {
  if num2 == 0 {
    return 0, errors.New("Division by zero")
  }

  return (float64(num1) / float64(num2)) * 100, nil
}

func updateStats(havePlayerWon bool) {
  gameModeStats.Played += 1
  if havePlayerWon {
    gameModeStats.Won += 1
  } 
}

func getDoorsLabels(doors []int) (labels []string) {
  for _,door := range doors {
    doorLabel := "Door " + strconv.Itoa(door)
    labels = append(labels, doorLabel) 
  }

  return
}

func getRandomDoor(doors []int) int {
  randomDoor := rand.Intn(len(doors))
  return doors[randomDoor]
}

func getRemainingDoor(door1 int, door2 int, doors []int) (remainingDoor int) {
  for _,door := range doors {
    if door!=door1 && door!=door2 {
      remainingDoor = door
      break
    }
  }
  return
}

func playGameMode() {
  doors := []int{1,2,3}
  winningDoor := getRandomDoor(doors)
  selectedDoorChoices := getDoorsLabels(doors) 
  selectedDoor := showMenu(selectedDoorChoices)
  emptyDoor := getRemainingDoor(selectedDoor, winningDoor, doors)
  remainingDoor := getRemainingDoor(selectedDoor, emptyDoor, doors)
  fmt.Printf("You have selected door %d. Door %d is empty. Do you want to change your choice to door %d?\n", selectedDoor, emptyDoor, remainingDoor)
  changeConfirmationChoices := []string{"Yes", "No"}
  changeConfirmation := showMenu(changeConfirmationChoices)

  switch changeConfirmation {
  case 1:
    selectedDoor = remainingDoor
    fmt.Printf("You have chosen to change your door. You have now door %d\n", selectedDoor)
  case 2:
    fmt.Printf("You have chosen to keep your door.\n")
  }
  var havePlayerWon bool 
  if selectedDoor == winningDoor {
    fmt.Println("Congratulations! You've just won a beautiful car.")
    havePlayerWon = true
  } else {
    fmt.Println("Oops! You lost. Take that stinky goat with you.")
  }
  updateStats(havePlayerWon)
  winningPercentage, _ := calculatePercentage(gameModeStats.Won, gameModeStats.Played) 
  fmt.Printf("Games played: %d | Games won: %d | Winning percentage: %.2f%%\n", gameModeStats.Played, gameModeStats.Won, winningPercentage)
}

func showMenu(choices []string) int {
	fmt.Println("\nEnter the number corresponding to your choice")

	for i := 0; i < len(choices); i++ {
		fmt.Println(i+1, "->", choices[i])
	}

	var userInput string

	fmt.Scanln(&userInput)
	userChoice, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Invalid input. Please retry!")

		return showMenu(choices)
	}

	if userChoice < 1 || userChoice > len(choices) {
		fmt.Println("Choice unavailable. Please retry!")

		return showMenu(choices)
	}
  
  return userChoice
}

func showMainMenu() int {
  mainMenu := []string{"Game mode", "Search mode", "Exit"}
  chosenMode := showMenu(mainMenu)

  return chosenMode
}

func main() {
	fmt.Println("Welcome to the monty hall game!")
  chosenMode := showMainMenu()

  for chosenMode != 3 {
    if chosenMode == 1 {
      playGameMode()
    } else {
      fmt.Println("Not implemented yet.")
    }
    chosenMode = showMainMenu()
  }

  fmt.Println("Bye! See you soon.")
}
