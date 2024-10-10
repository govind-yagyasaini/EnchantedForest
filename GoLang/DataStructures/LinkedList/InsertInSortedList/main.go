package main

import "fmt"

func main() {
	fmt.Println("Insertion in sorted Linked List.")
	myList := LinkedList{}
	myList.InsertAtEnd(10)
	myList.InsertAtEnd(20)
	myList.InsertAtEnd(30)
	myList.InsertAtEnd(50)
	myList.InsertAtEnd(60)

	myList.InsertInSortedList(40)
	myList.DisplayLinkedList()
	myList.InsertInSortedList(00)
	myList.DisplayLinkedList()
	myList.InsertInSortedList(70)
	myList.DisplayLinkedList()

}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insertion in sorted linked list
func (list *LinkedList) InsertInSortedList(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if list.head.data >= data {
		newNode.next = list.head
		list.head = newNode
		return
	}
	temp := list.head
	for temp.next != nil && temp.next.data < data {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode

}

// Insert nodes at the end
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode

}

// Display Linked List Nodes
func (list *LinkedList) DisplayLinkedList() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	temp := list.head
	for temp != nil {
		fmt.Printf("%d-> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")

}
