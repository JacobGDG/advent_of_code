package main

import (
  "fmt"
  "advent_of_code/input"
  "advent_of_code/days/three"
  "advent_of_code/days/four"
)

func main() {
  var token string
  var day int
  var task int

  fmt.Println("Enter your session token: ")
  fmt.Scanln(&token)

  fmt.Println("Enter the day to run: ")
  fmt.Scanln(&day)

  var inputData string = input.Get(token, day)

  fmt.Println("Enter the task to run (1 or 2): ")
  fmt.Scanln(&task)

  switch day {
  case 3:
    if task == 1 {
      three.TaskOne(inputData)
    } else {
      three.TaskTwo(inputData)
    }
  case 4:
    if task == 1 {
      four.TaskOne(inputData)
    } else {
      four.TaskOne(inputData)
    }
  }
}
