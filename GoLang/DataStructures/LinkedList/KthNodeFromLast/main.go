package main

import "fmt"

func main() {
	fmt.Println("Welcome to finding K-th node from last of a Linked List.")
	k := 5
	myList := LinkedList{}
	myList.InsertAtEnd(10)
	myList.InsertAtEnd(20)
	myList.InsertAtEnd(30)
	myList.InsertAtEnd(40)
	myList.InsertAtEnd(50)
	myList.DisplayLinkedList()
	myList.KthNodeFromLast(k)
	myList.FindKthNodeFromLast(k)
	myList.FindKthNodeFromEnd(k)
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Find k-th node from the end: Two-pass traversal approach O(n) Complexity
func (list *LinkedList) KthNodeFromLast(k int) {
	n := 1
	temp := list.head
	for temp.next != nil {
		temp = temp.next
		n++
	}

	if n < k {
		fmt.Println("There are fewer number of nodes in the list.")
		return

	}
	temp = list.head
	for i := 1; i < (n - k + 1); i++ {
		temp = temp.next

	}
	fmt.Printf("%d-th node from the end is: %d\n", k, temp.data)

}

// Find k-th node from the end using Hash Table- O(n) Complexity
func (list *LinkedList) FindKthNodeFromLast(k int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	flag := make(map[int]int)
	n := 1
	temp := list.head
	for temp != nil {
		flag[n] = temp.data
		temp = temp.next
		n++
	}
	if n-1 < k {
		fmt.Println("There are fewer number of nodes in the list.")
		return

	}
	fmt.Printf("%d-th node from the end is: %d\n", k, flag[n-k])

}

// Find k-th node from the end using two pointer approach- single pass solution- O(n) complexity
func (list *LinkedList) FindKthNodeFromEnd(k int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	if k <= 0 {
		fmt.Println("Invalid value for k.")
		return
	}

	first := list.head
	second := list.head

	// Move first pointer k steps ahead
	for i := 1; i < k; i++ {
		if first.next == nil {
			fmt.Printf("The list has fewer than %d elements.\n", k)
			return
		}
		first = first.next
	}

	for first.next != nil {
		first = first.next
		second = second.next
	}
	fmt.Printf("%d-th node from the end is: %d\n", k, second.data)
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
