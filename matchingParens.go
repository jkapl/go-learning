package main

import "fmt" 

var test = "()())("

func isValid (s string) bool {

  if len(s) % 2 != 0 {
    return false
  }
  
  stack := []byte{}
  
  for i := range s {
    fmt.Println(string(s[i]))
    if string(s[i]) == "(" {
      stack = append(stack, s[i])
    }
    if string(s[i]) == ")" {
      if len(stack) < 1 {
        return false
      }
      //current := stack[len(stack)-1]
      stack = stack[:len(stack)-1]
    }

  }
  if len(stack) > 0 {
    return false
  }
  return true
}

func main() {
  fmt.Println(isValid(test))
}