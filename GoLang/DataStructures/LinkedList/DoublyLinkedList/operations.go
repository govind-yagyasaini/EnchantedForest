package main

import (
	"fmt"
)

type DoublyNode struct {
	data int
	prev *DoublyNode
	next *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
}

// Insertion at beginning
func (list *DoublyLinkedList) InsertionAtBeginning(data int) {
	newNode := &DoublyNode{data: data}
	if list.head == nil {
		list.head = newNode
		newNode.prev = nil
		newNode.next = nil
		return

	}
	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode

}

// Insertion at end
func (list *DoublyLinkedList) InsertionAtEnd(data int) {
	newNode := &DoublyNode{data: data}
	if list.head == nil {
		list.head = newNode
		return

	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	newNode.prev = temp
	newNode.next = nil //Not required still I have written

}

// Insertion at specific position
func (list *DoublyLinkedList) InsertAtPosition(data int, position int) {
	newNode := &DoublyNode{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	if position == 1 {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
		return
	}

	temp := list.head
	for i := 1; i < position-1 && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		fmt.Println("Position out of bounds exception.")
		return
	}
	newNode.next = temp.next
	newNode.prev = temp
	if temp.next != nil {
		temp.next.prev = newNode
	}
	temp.next = newNode

}

// Deletion from beginning
func (list *DoublyLinkedList) DeletionFromBeginning() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return

	}
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	}

}

// Deletion from end
func (list *DoublyLinkedList) DeletionFromEnd() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if list.head.next == nil {
		list.head = nil
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next

	}
	temp.prev.next = nil

}

// Delete from position
func (list *DoublyLinkedList) DeletFromPosition(position int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if position == 1 {
		list.head = nil
		return
	}
	temp := list.head
	for i := 1; i < position-1 && temp != nil; i++ {
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("Position our of bounds exception.")
		return
	}
	if temp.next != nil {
		temp.next.prev = temp.prev
	}
	if temp.prev != nil {
		temp.prev.next = temp.next
	}

}

// Display Linked list values
func (list *DoublyLinkedList) DisplayLinkedList() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	temp := list.head
	for temp != nil {
		fmt.Printf("%d <-> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")

}
