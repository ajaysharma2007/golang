package list

import (
	"node"
)

// LinkedList Data Structure
type LinkedList struct {
	head *node.Node
}

// Add - adds data to the List
func (list *LinkedList) Add(data int) {
	node := new(node.Node)
	node.SetData(data)
	if list.head == nil {
		list.head = node
	} else {
		curr := list.head
		for curr.GetNext() != nil {
			curr = curr.GetNext()
		}
		curr.SetNext(node)
	}
}

// AddFront - adds data to the begining List
func (list *LinkedList) AddFront(data int) {
	node := new(node.Node)
	node.SetData(data)
	if list.head == nil {
		list.head = node
	} else {
		node.SetNext(list.head)
		list.head = node
	}
}

// RemoveFront - removes data from the begining of the list
func (list *LinkedList) RemoveFront() (node *node.Node) {
	if list.head == nil || list.head.GetNext() == nil {
		return list.head
	}
	returnNode := list.head
	list.head = list.head.GetNext()
	return returnNode
}

// Remove - removes data from the list
func (list *LinkedList) Remove() (node *node.Node) {
	if list.head == nil || list.head.GetNext() == nil {
		return list.head
	}
	curr := list.head
	for curr.GetNext().GetNext() != nil {
		curr = curr.GetNext()
	}
	returnNode := curr.GetNext()
	curr.SetNext(nil)
	return returnNode
}

// IsEmpty - return sa boolen to indiacte whether the list is empty of not.
func (list *LinkedList) IsEmpty() bool {
	if list.head == nil {
		return true
	}

	return false
}
