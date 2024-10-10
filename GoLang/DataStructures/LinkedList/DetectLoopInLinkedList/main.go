package main

import "fmt"

func main() {

	fmt.Println("Detecting loop in Linked List:")
	myList := LinkedList{}
	myList.InsertAtEnd(10)
	myList.InsertAtEnd(20)
	myList.InsertAtEnd(30)
	myList.InsertAtEnd(40)
	myList.InsertAtEnd(50)
	myList.InsertAtEnd(60)

	myList.head.next.next.next.next.next = myList.head //Create a loop
	myList.DetectLoop()
	myList.DetectLoopUsingHashSet()

}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Function to detect loop
// Floyd’s Cycle Detection Algorithm or Tortoise and Hare Algorithm.
// Time Complexity- O(n)
// Space Complexity- O(1)
func (list *LinkedList) DetectLoop() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}

	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next      //move 1 position
		fast = fast.next.next //move 2 position

		if slow == fast {
			fmt.Println("Loop Detected.")
			return
		}

	}
	fmt.Println("No Loop detected.")

}

// Detect Loops using Hash sets
// Time Complexity- O(n)
// Space Complexity- O(n)
// Less efficient than Tortoise Hare Algorithm

func (list *LinkedList) DetectLoopUsingHashSet() {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}
	visitedNodes := make(map[*Node]bool)
	temp := list.head

	for temp != nil {
		if _, found := visitedNodes[temp]; found {
			fmt.Println("Loop Detected.")
			return
		}

		visitedNodes[temp] = true
		temp = temp.next

	}
	fmt.Println("No Loop Detected.")

}

// Find loop beginning
func (list *LinkedList) FindLoopBeginning() *Node {

	if list.head == nil {
		fmt.Println("List is empty.")
		return nil
	}
	slow := list.head
	fast := list.head
	loopExists := false

	for fast != nil && fast.next != nil {
		slow = slow.next      //move 1 position
		fast = fast.next.next //move 2 position

		if slow == fast {
			fmt.Println("Loop Detected.")
			loopExists = true
			break
		}

	}

	if loopExists {
		slow = list.head
		for slow != fast {
			fast = fast.next
			slow = slow.next
		}
		return slow

	}
	fmt.Println("No Loop detected.")
	return nil

}

// Find Loop Length
func (list *LinkedList) FindLoopLength() int {

	if list.head == nil {
		fmt.Println("List is empty.")
		return 0
	}

	slow := list.head
	fast := list.head
	loopExists := false

	for fast != nil && fast.next != nil {
		slow = slow.next      //move 1 position
		fast = fast.next.next //move 2 position

		if slow == fast {
			fmt.Println("Loop Detected.")
			loopExists = true
			break
		}

	}
	counter := 1
	if loopExists {
		fast = fast.next //otherwise it will not enter the loop
		for slow != fast {
			fast = fast.next
			counter++
		}
		return counter

	}
	fmt.Println("No Loop detected.")
	return 0

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
