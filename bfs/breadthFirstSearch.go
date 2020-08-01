package main

import "fmt"

type Node struct {
  Val int
  left *Node
  right *Node
}

func traverse (tree Node) {
  queue := []Node{tree}
  for ;len(queue)>0; {
    current := queue[0]
    fmt.Println(current.Val)
    
    if current.left != nil {
      queue = append(queue, *current.left)
    }
    if current.right != nil {
      queue = append(queue, *current.right)
    }
    queue = queue[1:]
  }
}

func main() {
  root := Node{1, nil, nil}
  left := Node{2, nil, nil}
  right := Node{3, nil, nil}
  root.left = &left
  root.right = &right
  traverse(root)
}