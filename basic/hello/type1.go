package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	val int
}

func main() {
	fmt.Println("hello world with a struct")
	// new struct decl.
	var foo = Node{nil, nil, 42}
	fmt.Println(foo.val)
}