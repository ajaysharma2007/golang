package main

import "fmt"

var numOp int
var size int
var noOfFriends int

// WeightedQuickUnion type
type WeightedQuickUnion struct {
	rootArr       []int
	sizeArr       []int
	numComponents int
}

func newWeightedQuickUnion(n int) *WeightedQuickUnion {
	newRootArr := make([]int, n, n)
	newSizeArr := make([]int, n, n)
	for i := 0; i < n; i++ {
		newRootArr[i] = i
		newSizeArr[i] = 1
	}
	noOfFriends = n
	return &WeightedQuickUnion{newRootArr, newSizeArr, n}
}

func (qu *WeightedQuickUnion) getRoot(p int) int {
	i := p
	for qu.rootArr[i] != i {
		qu.rootArr[i] = qu.rootArr[qu.rootArr[i]]
		i = qu.rootArr[i]
	}
	return i
}

func (qu *WeightedQuickUnion) allConnected() bool {
	return size == noOfFriends
}

func (qu *WeightedQuickUnion) allConnectedComponents() bool {
	return qu.numComponents == 1
}

func (qu *WeightedQuickUnion) union(p int, q int) {

	if !qu.isConnected(p, q) {
		numOp++
		qu.numComponents--
		rootP := qu.getRoot(p)
		rootQ := qu.getRoot(q)

		if qu.sizeArr[p] < qu.sizeArr[q] {
			qu.rootArr[rootP] = qu.rootArr[rootQ]
			size = qu.sizeArr[rootQ] + qu.sizeArr[rootP]
			qu.sizeArr[rootQ] = size
		} else {
			qu.rootArr[rootQ] = qu.rootArr[rootP]
			size = qu.sizeArr[rootP] + qu.sizeArr[rootQ]
			qu.sizeArr[rootP] = size
		}
	}
}

func (qu *WeightedQuickUnion) isConnected(p int, q int) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}

func main() {
	n := 10
	var uf = newWeightedQuickUnion(n)

	for i := 0; i < n-1; i++ {
		uf.union(i, i+1)
		fmt.Println(uf.allConnected(), numOp)
		fmt.Println(uf.allConnectedComponents(), numOp)
	}
}
