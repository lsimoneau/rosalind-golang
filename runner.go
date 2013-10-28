package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

func main() {
  input, err := ReadLines("input.txt")
  dna := strings.Join(input[1:], "")
  if err == nil {
    fmt.Println(Revp(dna))
  }
}

func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
