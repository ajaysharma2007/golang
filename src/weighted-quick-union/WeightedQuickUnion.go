package main

import "fmt"

// WeightedQuickUnion type
type WeightedQuickUnion struct {
	rootArr []int
	sizeArr []int
}

func newWeightedQuickUnion(n int) *WeightedQuickUnion {
	newRootArr := make([]int, n, n)
	newSizeArr := make([]int, n, n)
	for i := 0; i < n; i++ {
		newRootArr[i] = i
		newSizeArr[i] = 1
	}
	return &WeightedQuickUnion{newRootArr, newSizeArr}
}

func (qu *WeightedQuickUnion) getRoot(p int) int {
	i := p
	for qu.rootArr[i] != i {
		qu.rootArr[i] = qu.rootArr[qu.rootArr[i]]
		i = qu.rootArr[i]
	}
	return i
}

func (qu *WeightedQuickUnion) union(p int, q int) {
	rootP := qu.getRoot(p)
	rootQ := qu.getRoot(q)

	if qu.sizeArr[p] < qu.sizeArr[q] {
		qu.rootArr[rootP] = qu.rootArr[rootQ]
		qu.sizeArr[rootQ] = qu.sizeArr[rootQ] + qu.sizeArr[rootP]
	} else {
		qu.rootArr[rootQ] = qu.rootArr[rootP]
		qu.sizeArr[rootP] = qu.sizeArr[rootP] + qu.sizeArr[rootQ]
	}
}

func (qu *WeightedQuickUnion) isConnected(p int, q int) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}

func main() {
	var uf = newWeightedQuickUnion(5)
	uf.union(0, 2)
	uf.union(1, 2)
	fmt.Println(uf.isConnected(0, 2))
}
