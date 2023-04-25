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

func getProposableDoor(selectedDoor int, winningDoor int, doors []int) (proposableDoor int) {
  for _,door := range doors {
    if door!=selectedDoor && door!=winningDoor {
      proposableDoor = door
      break
    }
  }
  return
}

func getRemainingDoors(selectedDoor int, doors []int) []int {
  var remainingDoors []int
  for _,door := range doors {
    if door != selectedDoor {
      remainingDoors = append(remainingDoors, door)
    }
  }
  return remainingDoors
}

func playGameMode() {
  doors := []int{1,2,3}
  winningDoor := getRandomDoor(doors)
  selectedDoorChoices := getDoorsLabels(doors) 
  selectedDoor := showMenu(selectedDoorChoices)
  remainingDoors := getRemainingDoors(selectedDoor, doors) 
  doorToPropose := getProposableDoor(selectedDoor, winningDoor, doors)
  fmt.Printf("You have selected door %d. Door %d is empty. Do you want to change your choice?\n", selectedDoor, doorToPropose)
  changeConfirmationChoices := []string{"Yes", "No"}
  changeConfirmation := showMenu(changeConfirmationChoices)

  switch changeConfirmation {
  case 1:
    fmt.Println("You have chosen to change your door. Please select a new door ?")
    remainingDoorsLabels := getDoorsLabels(remainingDoors)
    newDoor := showMenu(remainingDoorsLabels)
    fmt.Printf("You've changed your door. You have now door %d\n", newDoor)
  case 2:
    fmt.Printf("You have chosen to keep your door.\n")
  }

}

func showMenu(choices []string) int {
	fmt.Println("Enter the number corresponding to your choice")

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

func main() {
	fmt.Println("Welcome to the monty hall game!")
  mainMenu := []string{"Game mode", "Search mode", "Exit"}
  chosenMode := showMenu(mainMenu)
  if chosenMode == 1 {
    playGameMode()
  }
}
