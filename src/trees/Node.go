package tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (node *Node) SetData(data int) {
	node.data = data
}

func (node *Node) GetData() int {
	return node.data
}

func (node *Node) SetLeft(left *Node) {
	node.left = left
}

func (node *Node) SetRight(right *Node) {
	node.right = right
}

func (node *Node) GetRight() *Node {
	return node.right
}

func (node *Node) GetLeft() *Node {
	return node.left
}
