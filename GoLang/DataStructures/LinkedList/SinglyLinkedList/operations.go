package main

import "fmt"

type Node struct {
	data int   //data field
	next *Node //pointer to next node
}

type LinkedList struct {
	head *Node //Pointer to the first node
}

// Insert new node at the beginning
func (list *LinkedList) InsertAtBeginning(data int) {
	//Create a new node
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

// Insert at the end
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		fmt.Println("Linked List is empty")
		list.head = newNode
		return

	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode

}

//Insert at any position

func (list *LinkedList) InsertAtPosition(data int, position int) {
	newNode := &Node{data: data}
	if position == 1 {
		newNode.next = list.head
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
	temp.next = newNode

}

// Delete first node
func (list *LinkedList) DeleteFirstNode() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	list.head = list.head.next

}

// Delete from any position
func (list *LinkedList) DeleteFromPosition(position int) {

	if position == 1 {
		fmt.Println("You are removing first node")
		list.head = list.head.next
		return
	}
	temp := list.head
	for i := 1; i < position-1 && temp != nil; i++ {
		temp = temp.next
	}
	if temp == nil || temp.next == nil {
		fmt.Println("Position out of bounds exception.")
		return
	}
	temp.next = temp.next.next

}

// Delete from last
func (list *LinkedList) DeleteFromLast() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return

	}

	if list.head.next == nil {
		fmt.Println("List has only one node.")
		list.head = nil
		return

	}
	temp := list.head
	for temp.next.next != nil {
		temp = temp.next

	}
	temp.next = nil

}

// Display Linked List
func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("Linked List is empty.")
		return
	}
	temp := list.head
	for temp != nil {
		fmt.Printf("%d-> ", temp.data)
		temp = temp.next

	}
	fmt.Println("NULL")
}
