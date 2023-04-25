package main

import (
	"fmt"
	"strconv"
  "math/rand"
)

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
  
  if selectedDoor == winningDoor {
    fmt.Println("Congratulations! You've just won a beautiful car.")
  } else {
    fmt.Println("Oops! You lost. Take that stinky goat with you.")
  }

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
