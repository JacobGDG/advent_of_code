package main

import (
  "fmt"
  "advent_of_code/getinput"
)

func main() {
  fmt.Println("Enter your session token: ")

  var token string

  fmt.Scanln(&token)

  fmt.Println(getinput.GetInput(token))
}
