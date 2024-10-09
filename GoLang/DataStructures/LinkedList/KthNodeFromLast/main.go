package main

import "fmt"

func main() {
	fmt.Println("Welcome to finding K-th node from last of a Linked List.")
	k := 1
	myList := LinkedList{}
	myList.InsertAtEnd(10)
	myList.InsertAtEnd(20)
	myList.InsertAtEnd(30)
	myList.InsertAtEnd(40)
	myList.InsertAtEnd(50)
	myList.DisplayLinkedList()
	data := myList.KthNodeFromLast(k)
	fmt.Printf("%d-th node from the end is: %d\n", k, data.data)
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Find k-th node from the end
func (list *LinkedList) KthNodeFromLast(k int) *Node {

	n := 1
	temp := list.head
	for temp.next != nil {
		temp = temp.next
		n++
	}

	if n < k-1 {
		fmt.Println("There are fewer number of nodes in the list.")
		return nil

	}
	temp = list.head
	for i := 1; i < (n - k + 1); i++ {
		temp = temp.next

	}
	return temp

}

// Insert nodes from end of a linked list
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
	newNode.next = nil

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
