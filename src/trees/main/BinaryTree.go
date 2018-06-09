package main

import (
	"fmt"
	"trees"
)

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
		fmt.Println(node)

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

func main() {
	arr := [3]int{3, 1, 4}
	x := arr[0:3]
	binary := BinaryTree{nil, x, 0}
	for _, x := range binary.arr {
		binary.add(x)
	}
	binary.breadthFirst()
}
