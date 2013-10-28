package main

import (
  "strings"
  "fmt"
  "bytes"
)

func Revp(dna string) string {
  var buffer bytes.Buffer

  for i := 0; i < len(dna); i++ {
    found, length := siteAtLocation(dna, i)

    if found {
      buffer.WriteString(fmt.Sprintf("%v %v\n", i + 1, length))
    }
  }

  return buffer.String()
}

func siteAtLocation(dna string, location int) (isSite bool, length int) {
  for i := 12; i >= 4; i-- {
    j := min(i, len(dna) - location)

    if j >= 4 {
      toCheck := dna[location:location + j]
      if rComp(toCheck) == reverse(toCheck) {
        return true, j
      }
    }
  }
  return false, 0
}

func rComp(dna string) string {
  return strings.Map(comp, dna)
}

func comp(r rune) rune {
  switch {
    case r == 'A':
      return 'T'
    case r == 'T':
      return 'A'
    case r == 'G':
      return 'C'
    case r == 'C':
      return 'G'
  }
  return r
}

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func min(x int, y int) int {
  if x > y {
    return y
  } else {
    return x
  }
}
