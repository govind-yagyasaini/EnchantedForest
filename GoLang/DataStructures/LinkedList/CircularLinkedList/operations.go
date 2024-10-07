package main

import "fmt"

type CircularNode struct {
	data int
	next *CircularNode
}

type CircularLinkedList struct {
	head *CircularNode
}

// Insertion at beginning
func (list *CircularLinkedList) InsertAtBeginning(data int) {
	newNode := &CircularNode{data: data}
	if list.head == nil {
		list.head = newNode
		list.head.next = list.head
		return
	}
	temp := list.head
	for temp.next != list.head {
		temp = temp.next
	}
	newNode.next = list.head
	temp.next = newNode
	list.head = newNode

}

// Insert at End
func (list CircularLinkedList) InsertAtEnd(data int) {
	newNode := &CircularNode{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
		return
	}
	temp := list.head
	for temp.next != list.head {
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = list.head
}

// Insert at position
func (list *CircularLinkedList) InsertAtPosition(data int, position int) {
	newNode := &CircularNode{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
		return
	}
	if position == 1 {
		temp := list.head
		// Find the last node
		for temp.next != list.head {
			temp = temp.next
		}
		newNode.next = list.head
		temp.next = newNode // Update the last node to point to new node
		list.head = newNode // Update head to point to the new node
		return
	}
	temp := list.head
	for i := 1; i < position-1 && temp.next != list.head; i++ {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode

}

// Delete from beginning
func (list *CircularLinkedList) DeletionFromBeginning() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}

	if list.head.next == list.head {
		list.head = nil
		return
	}
	temp := list.head
	for temp.next != list.head {
		temp = temp.next
	}
	list.head = list.head.next
	temp.next = list.head

}

// Deletion from end
func (list *CircularLinkedList) DeletionFromEnd() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if list.head.next == list.head {
		list.head = nil
		return
	}
	temp := list.head
	for temp.next.next != list.head {
		temp = temp.next
	}
	temp.next = list.head

}

// Deletion from any position
func (list *CircularLinkedList) DeleteFromPosition(position int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return

	}
	temp := list.head
	if position == 1 {
		if list.head.next == list.head { // Only one node
			list.head = nil
			return
		}
		for temp.next != list.head {
			temp = temp.next
		}
		list.head = list.head.next
		temp.next = list.head
		return
	}

	for i := 1; i < position-1 && temp.next != list.head; i++ {
		temp = temp.next

	}

	if temp.next == list.head || temp.next == nil {
		fmt.Println("Position out of bounds exception.")
		return
	}
	temp.next = temp.next.next

}

// Display Circular Linked List
func (list *CircularLinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	temp := list.head
	for {
		fmt.Printf("%d ->", temp.data)
		temp = temp.next
		if temp == list.head {
			break
		}

	}
	fmt.Println("(head)")

}
