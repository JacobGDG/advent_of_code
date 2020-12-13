package main

import (
  "fmt"
  "advent_of_code/input"
  "advent_of_code/days"
)

func main() {
  var token string
  var day int
  var task int

  fmt.Println("Enter your session token: ")
  fmt.Scanln(&token)

  fmt.Println("Enter the day to run: ")
  fmt.Scanln(&day)

  var input_data string = input.Get(token, day)

  fmt.Println("Enter the task to run (1 or 2): ")
  fmt.Scanln(&task)

  switch day {
  case 3:
    if task == 1 {
      three.TaskOne(input_data)
    } else {
      three.TaskTwo(input_data)
    }
  }
}
