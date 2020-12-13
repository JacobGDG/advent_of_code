package three

import (
  "fmt"
  "strings"
)

type Grid struct {
  matrix [][]string
  depth int
  width int
}

func createGrid(rawData string) Grid {
  lines := strings.Split(rawData, "\n")
  lineNum := len(lines) - 1

  matrix := make([][]string, lineNum)

  for i := 0; i < lineNum; i++ {
    matrix[i] = strings.Split(lines[i], "")
  }

  grid := Grid{
    matrix: matrix,
    depth: len(matrix),
    width: len(matrix[0]),
  }

  return grid
}

func traverseHill(grid Grid, across int, down int) int {
  totalTrees := 0

  x := 0
  for y := 0; y < grid.depth; y += down {
    if y == 0 {
      continue
    }

    x += across
    if x >= grid.width {
      x -= grid.width
    }

    if grid.matrix[y][x] == "#" {
      totalTrees ++
    }
  }

  return totalTrees
}

func TaskOne(input string) {
  grid := createGrid(input)

  fmt.Println(traverseHill(grid, 3, 1))
}

func TaskTwo(input string) {
  grid := createGrid(input)
  answer := 1
  var num int

  num = traverseHill(grid, 1, 1)
  answer *= num

  num = traverseHill(grid, 3, 1)
  answer *= num

  num = traverseHill(grid, 5, 1)
  answer *= num

  num = traverseHill(grid, 7, 1)
  answer *= num

  num = traverseHill(grid, 1, 2)
  answer *= num

  fmt.Println(answer)
}

