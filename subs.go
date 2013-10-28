package main

import (
  "strings"
)

func Subs(needle string, haystack string) []int {
  c := make(chan int)
  var result []int

  go findMatches(needle, haystack, c)

  for i := range c {
    result = appendIfMissing(result, i)
  }

  return result
}

func findMatches(needle string, haystack string, c chan int) {
  for i := 0; i < len(haystack) - len(needle); i++ {
    loc := strings.Index(haystack[i:], needle)

    if loc != -1 {
      c <- loc + i + 1
    }
  }

  close(c)
}

func appendIfMissing(slice []int, i int) []int {
  for _, ele := range slice {
    if ele == i {
      return slice
    }
  }
  return append(slice, i)
}
