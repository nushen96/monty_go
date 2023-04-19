package main

import (
  "fmt"
  "strconv"
)


func show_menu(choices []string) {
  fmt.Println("Enter the number corresponding to your choice")
  for i:=0; i<len(choices); i++ {
    fmt.Println(i+1,"->",choices[i])
  }
  var user_input string
  fmt.Scanln(&user_input)
  user_choice, err := strconv.Atoi(user_input)

  if err != nil {
    fmt.Println("Invalid input. Please retry!")
    show_menu(choices)
    return
  }

  if user_choice<1 || user_choice>len(choices) {
    fmt.Println("Choice unavailable. Please retry!")
    show_menu(choices)
    return
  }

  fmt.Println(user_choice)
  fmt.Println("You chose to", choices[user_choice-1])
}


func main() {
  fmt.Println("Welcome to the monty hall game!")
  show_menu([]string{"Play", "Quit"})
}
