package stack

import (
	"errors"
	"node"
)

// Stack - stack data structure
type Stack struct {
	top  *node.Node
	size int
}

// Push - adds data to the top of the stack
func (stack *Stack) Push(data int) {
	node := new(node.Node)
	node.SetData(data)
	node.SetNext(stack.top)
	stack.top = node
}

// Pop - removes data from the top
func (stack *Stack) Pop() (returnNode *node.Node, err error) {
	if stack.top == nil {
		err = errors.New("Can't pop from empty stack")
	}
	returnNode = stack.top
	stack.top = stack.top.GetNext()
	return
}

//Peek - gives view of the top most element without removing it
func (stack *Stack) Peek() (returnNode *node.Node, err error) {
	if stack.top == nil {
		err = errors.New("Can't peek from empty stack")
	}

	return stack.top, err
}

// IsEmpty - returns a bool indiacting whether a stack is empty or not
func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}
