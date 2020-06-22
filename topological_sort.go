package main

import "fmt"

var input = map[rune][]rune{
  'a': []rune{'b','c'},
  'b': []rune{'d'},
  'c': []rune{},
  'd': []rune{'e'},
  'e': []rune{},
}

func createDigraph (input map[rune][]rune) map[rune]int {
  var output = make(map[rune]int)

  for key, _ := range input {
    output[key] = 0
  }

  for _, value := range input {
      for _, v := range value {
        output[v] += 1
      }
  }
  return output
}

func sort (input map[rune][]rune) []rune {
  nodes_with_no_incoming_edges := make([]rune, 0)
  order := make([]rune, 0)
  node := 'x'

  digraph := createDigraph(input)
  fmt.Println(digraph)
  
  for key, value := range digraph {
    if value == 0 {
      nodes_with_no_incoming_edges = append(nodes_with_no_incoming_edges, key)
    }
  }

  for len(nodes_with_no_incoming_edges) > 0 {

    node, nodes_with_no_incoming_edges = nodes_with_no_incoming_edges[len(nodes_with_no_incoming_edges)-1], nodes_with_no_incoming_edges[:len(nodes_with_no_incoming_edges)-1]

    order = append(order, node)

    for _, v := range input[node] {
      digraph[v] -= 1
      if digraph[v] == 0 {
        nodes_with_no_incoming_edges = append(nodes_with_no_incoming_edges, v)
      }
    }
  }


  return order
}

func main() {
  output := sort(input)
  for _, v := range output {
    fmt.Printf("%c \n", v)
  }
}