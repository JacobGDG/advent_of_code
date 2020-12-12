package main

import (
  "fmt"
  "advent_of_code/getinput"
)

func main() {
  var token string
  var day int

  fmt.Println("Enter your session token: ")
  fmt.Scanln(&token)

  fmt.Println("Enter the task to run: ")
  fmt.Scanln(&day)

  fmt.Println(getinput.GetInput(token, day))
}
