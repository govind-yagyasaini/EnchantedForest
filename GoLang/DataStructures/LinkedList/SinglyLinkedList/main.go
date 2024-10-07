package main

import "fmt"

func main() {

	fmt.Println("Welcome to Linked List in GoLang")

	node1 := &Node{data: 20}
	node2 := &Node{data: 30}
	node3 := &Node{data: 40}

	node1.next = node2
	node2.next = node3
	node3.next = nil

	myList := LinkedList{head: node1}

	myList.InsertAtBeginning(10)
	myList.InsertAtEnd(50)
	myList.InsertAtPosition(60, 6)
	myList.DeleteFirstNode()
	myList.DeleteFromPosition(2)
	myList.DeleteFromPosition(8)
	myList.DeleteFromLast()
	myList.Display()

}
