package main

import (
	"errors"
	"fmt"
	"math"
	"trees"
)

// BinaryTree - Binary tree data type. Stores ptr to root node
type BinaryTree struct {
	root *tree.Node
	arr  []int
	size int
}

func (binarytree *BinaryTree) breadthFirst() {
	if binarytree.root == nil {
		return
	}

	queue := tree.Queue{-1, -1, [3]*tree.Node{}}

	queue.Enqueue(binarytree.root)

	for queue.IsEmpty() == false {
		node, err := queue.Dequeue()

		if err == nil {
			left := node.GetLeft()
			if left != nil {
				queue.Enqueue(left)
			}
			right := node.GetRight()
			if right != nil {
				queue.Enqueue(right)
			}
		}
	}
}

func (binarytree *BinaryTree) lookup(node *tree.Node, data int) (*tree.Node, error) {
	if node == nil {
		return nil, errors.New("Node not found")
	}

	if data == node.GetData() {
		return node, nil
	}

	if data < node.GetData() {
		return binarytree.lookup(node.GetLeft(), data)
	}

	return binarytree.lookup(node.GetRight(), data)
}

func (binarytree *BinaryTree) maxDepth(node *tree.Node, count int) int {
	if node == nil {
		return count
	}
	count++
	leftCount := binarytree.maxDepth(node.GetLeft(), count)
	rightCount := binarytree.maxDepth(node.GetRight(), count)

	return int(math.Max(float64(leftCount), float64(rightCount)))
}

func (binarytree *BinaryTree) findMinimum(node *tree.Node) *tree.Node {
	if node == nil {
		return node
	}
	if node.GetLeft() == nil {
		return node
	}
	return binarytree.findMinimum(node.GetLeft())
}

func (binarytree *BinaryTree) addRecursive(node *tree.Node, data int) *tree.Node {
	insert := new(tree.Node)
	insert.SetData(data)
	if node == nil {
		return node
	}
	if data < binarytree.root.GetData() {
		binarytree.root.SetLeft(binarytree.addRecursive(binarytree.root.GetLeft(), data))
	} else {
		binarytree.root.SetRight(binarytree.addRecursive(binarytree.root.GetRight(), data))
	}
	return node
}

func (binarytree *BinaryTree) add(data int) {
	node := new(tree.Node)
	node.SetData(data)

	if binarytree.size == 0 {
		binarytree.root = node
		binarytree.size++
		return
	}

	currNode := binarytree.root
	for currNode != nil {
		if data < currNode.GetData() {
			if currNode.GetLeft() == nil {
				currNode.SetLeft(node)
				binarytree.size++
				return
			}
			currNode = currNode.GetLeft()
		} else {
			if currNode.GetRight() == nil {
				currNode.SetRight(node)
				binarytree.size++
				return
			}
			currNode = currNode.GetRight()
		}
	}
}

// PreOrderTraverse - Traverses tree in pre order.``
func (binarytree *BinaryTree) PreOrderTraverse(node *tree.Node) {
	if node == nil {
		return
	}

	binarytree.PreOrderTraverse(node.GetLeft())
	binarytree.PreOrderTraverse(node.GetRight())
}

// InOrderTraverse - Traverses tree in order.
func (binarytree *BinaryTree) InOrderTraverse(node *tree.Node) {
	if node == nil {
		return
	}

	binarytree.PreOrderTraverse(node.GetLeft())
	binarytree.PreOrderTraverse(node.GetRight())
}

// PostOrderTraverse - Traverses tree post order.
func (binarytree *BinaryTree) PostOrderTraverse(node *tree.Node) {
	if node == nil {
		return
	}

	binarytree.PreOrderTraverse(node.GetLeft())
	binarytree.PreOrderTraverse(node.GetRight())
}

func main() {
	arr := [3]int{1, 2, 3}
	x := arr[0:3]
	binary := BinaryTree{nil, x, 0}
	for _, x := range binary.arr {
		binary.add(x)
	}
	fmt.Println(binary.maxDepth(binary.root, -1))
}
