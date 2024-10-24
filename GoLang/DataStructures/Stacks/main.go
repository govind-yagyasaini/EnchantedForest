package main

import "fmt"

func main() {
	fmt.Println("Welcome to Stacks in GoLang")
	stack := Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	fmt.Println("Elements after inserting values in stack-")
	stack.Display()
	top, _ := stack.Pop()
	fmt.Println("Top most element is:", top)
	fmt.Println("Elements after pop operation in stack-")
	stack.Display()
	top = stack.Peek()
	fmt.Println("Top most element is:", top)
	fmt.Println("Elements after peek operation in stack-")
	stack.Display()

}
