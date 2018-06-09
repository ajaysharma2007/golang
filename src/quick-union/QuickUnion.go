package main

import (
	"fmt"
)

// QuickUnion type
type QuickUnion struct {
	arr []int
}

func newQuickUnion(n int) *QuickUnion {
	newArr := make([]int, n, n)
	for i := 0; i < n; i++ {
		newArr[i] = i
	}

	return &QuickUnion{newArr}
}

func (qu *QuickUnion) getRoot(p int) int {
	i := p
	for qu.arr[i] != i {
		i = qu.arr[i]
	}

	return i
}

func (qu *QuickUnion) union(p int, q int) {
	rootP := qu.getRoot(p)
	rootQ := qu.getRoot(q)

	qu.arr[rootP] = qu.arr[rootQ]
}

func (qu *QuickUnion) isConnected(p int, q int) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}

func main() {
	var uf = newQuickUnion(5)
	uf.union(0, 2)
	uf.union(1, 2)
	fmt.Println(uf.isConnected(0, 2))
}
