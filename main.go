package main

import (
  "fmt"
  "advent_of_code/input"
)

func main() {
  var token string
  var day int

  fmt.Println("Enter your session token: ")
  fmt.Scanln(&token)

  fmt.Println("Enter the task to run: ")
  fmt.Scanln(&day)

  fmt.Println(input.Get(token, day))
}
