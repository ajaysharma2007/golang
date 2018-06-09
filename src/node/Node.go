package node

// Node data type
type Node struct {
	data int
	next *Node
}

// SetData - sets the data to the node
func (node *Node) SetData(data int) {
	node.data = data
}

// GetData = gets node data
func (node *Node) GetData() int {
	return node.data
}

// GetNext - retrieves the next node
func (node *Node) GetNext() *Node {
	return node.next
}

// SetNext - sets next node for the current node
func (node *Node) SetNext(nextNode *Node) {
	node.next = nextNode
}
