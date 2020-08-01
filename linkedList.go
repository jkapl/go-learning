package main

import "fmt"

type node struct {
  value int
  next *node
}


func main() {
  head := node{5, nil}
  tail := node{6,nil}
  head.next = &tail
  fmt.Println(head)
}