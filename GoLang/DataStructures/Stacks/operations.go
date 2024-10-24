package main

import "fmt"

// Define the stack structure
type Stack struct {
	elements []int
}

// Pushes an element on top of the stack
func (s *Stack) Push(data int) {
	s.elements = append(s.elements, data)

}

// Removes the top most element from stack
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty.")
		return 0, false
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true

}

// Returns the to most element without deleting it
func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty.")
		return 0
	}
	top := s.elements[len(s.elements)-1]
	return top

}

// Displays all elements of a stack
func (s *Stack) Display() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Println("Elements are:", s.elements)
		//fmt.Printf("Length of the stack is %d\n", len(s.elements))
	}

}
