// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head  Node
	nodes []Node
}

func main() {
	secondNode := Node{data: 1, next: nil}
	headNode := Node{data: 0, next: &secondNode}
	linkedList := LinkedList{
		head:  headNode,
		nodes: []Node{headNode, secondNode},
	}
	fmt.Println(linkedList.head)
}
