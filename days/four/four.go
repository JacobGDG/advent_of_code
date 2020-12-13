package four

import (
  "fmt"
  "strings"
  "strconv"
  "regexp"
)

type document struct {
  byr int
  iyr int
  eyr int
  hgt string
  hcl string
  ecl string
  pid string
  cid string
}
func (d document) basicValid() bool {
  if d.byr == 0 {
    return false
  }
  if d.iyr == 0 {
    return false
  }
  if d.eyr == 0 {
    return false
  }
  if d.hgt == "" {
    return false
  }
  if d.hcl == "" {
    return false
  }
  if d.ecl == "" {
    return false
  }
  if d.pid == "" {
    return false
  }

  return true
}
func (d document) fullValid() bool {
  if d.byr == 0 || d.byr < 1920 || d.byr > 2002 {
    return false
  }
  if d.iyr == 0 || d.iyr < 2010 || d.iyr > 2020 {
    return false
  }
  if d.eyr == 0 || d.eyr < 2020 || d.eyr > 2030 {
    return false
  }
  if d.hgt == "" {
    return false
  }
  if regexp.MatchString("\\d+(cm|in)", d.hgt) {
  } else {
    return false
  }
  if d.hcl == "" {
    return false
  }
  if d.ecl == "" {
    return false
  }
  if d.pid == "" {
    return false
  }

  return true
}

func createDoc(docString string) document {
  keyValues := strings.Split(docString, " ")
  doc := document{}

  for i := 0; i < len(keyValues); i++ {
    keyValue := strings.Split(keyValues[i], ":")

    switch keyValue[0] {
      case "byr":
        doc.byr, _ = strconv.Atoi(keyValue[1])
      case "iyr":
        doc.iyr, _ = strconv.Atoi(keyValue[1])
      case "eyr":
        doc.eyr, _ = strconv.Atoi(keyValue[1])
      case "hgt":
        doc.hgt = keyValue[1]
      case "hcl":
        doc.hcl = keyValue[1]
      case "ecl":
        doc.ecl = keyValue[1]
      case "pid":
        doc.pid = keyValue[1]
      case "cid":
        doc.cid = keyValue[1]
    }
  }

  return doc
}

func sortDocuments(rawInput string) []document {
  lines := strings.Split(rawInput, "\n")
  lineNum := len(lines)

  var docs []document
  currentRawDoc := ""

  for i := 0; i < lineNum; i++ {
    line := lines[i]
    if line == "" {
      docs = append(docs, createDoc(currentRawDoc))
      currentRawDoc = ""
      continue
    }
    currentRawDoc = currentRawDoc + " " + line
  }

  return docs
}

func TaskOne(input string) {
  documents := sortDocuments(input)
  validCount := 0

  for i := 0; i < len(documents); i++ {
    if documents[i].basicValid() {
      validCount++
    }
  }

  fmt.Println(validCount)
}

func TaskTwo(input string) {
  documents := sortDocuments(input)
  validCount := 0

  for i := 0; i < len(documents); i++ {
    if documents[i].fullValid() {
      validCount++
    }
  }

  fmt.Println(validCount)
}
