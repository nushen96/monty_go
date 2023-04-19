package main

import (
  "fmt"
  "strconv"
)

func is_num(s string) bool {
  _, err := strconv.Atoi(s)
  return err == nil
}

func show_menu(choices []string) {
  fmt.Println("Enter the number corresponding to your choice")
  for i:=0; i<len(choices); i++ {
    fmt.Println(i+1,"->",choices[i])
  }
  var user_input string
  fmt.Scanln(&user_input)
  fmt.Println("You entered", user_input)
  if !is_num(user_input) {
    fmt.Println("Wrong choice. Please retry!")
    show_menu(choices)
  }
  user_choice, _ := strconv.Atoi(user_input)
  if user_choice<1 || user_choice>len(choices) {
    fmt.Println("Wrong choice. Please retry!")
    show_menu(choices)
  }
  fmt.Println(user_choice)
  fmt.Println("You chose to", choices[user_choice-1])
}


func main() {
  fmt.Println("Welcome to the monty hall game!")
  show_menu([]string{"Play", "Quit"})
}
