package four

import (
  "fmt"
  "strings"
)

func sortDocuments(rawInput string) {
  lines := strings.Split(rawInput, "\n")
  lineNum := len(lines)

  var rawDocs []string
  currentRawDoc := ""

  for i := 0; i < lineNum; i++ {
    line := lines[i]
    if line == "" {
      rawDocs = append(rawDocs, currentRawDoc)
      currentRawDoc = ""
      continue
    }
    currentRawDoc = currentRawDoc + " " + line
  }
}

func TaskOne(input string) {
  sortDocuments(input)
}
