package main

import "fmt"

func main() {
	fmt.Println("Welcome to Linked Lists in GoLang.")

	node1 := &DoublyNode{data: 20}
	node2 := &DoublyNode{data: 30}
	node3 := &DoublyNode{data: 40}
	node4 := &DoublyNode{data: 50}

	node1.prev = nil
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2
	node3.next = node4
	node4.prev = node3
	node4.next = nil

	myList := DoublyLinkedList{head: node1}
	myList.InsertionAtBeginning(10)
	myList.InsertionAtEnd(60)
	myList.InsertAtPosition(70, 7)
	myList.InsertAtPosition(333, 333)
	myList.DeletionFromBeginning()
	myList.DeletionFromEnd()
	myList.DeletFromPosition(6)
	myList.DisplayLinkedList()

}
