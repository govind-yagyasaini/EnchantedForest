package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Circular Linked List in GoLang")
	node1 := &CircularNode{data: 10}
	node2 := &CircularNode{data: 20}
	node3 := &CircularNode{data: 30}

	node1.next = node2
	node2.next = node3
	node3.next = node1

	myList := CircularLinkedList{head: node1}
	myList.InsertAtBeginning(00)
	myList.InsertAtEnd(40)
	myList.InsertAtPosition(50, 6)
	myList.DeleteFromPosition(3)
	myList.DeletionFromBeginning()
	myList.DeletionFromEnd()
	myList.Display()

}
