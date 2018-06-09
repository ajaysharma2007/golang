package main

import "fmt"

// QuickFInd type
type QuickFInd struct {
	arr []int
}

// QuickFIndNew nion Object
func QuickFIndNew(n int) *QuickFInd {

	newArr := make([]int, n, n)

	for i := 0; i < n; i++ {
		newArr[i] = i
	}
	return &QuickFInd{newArr}
}

func (uf *QuickFInd) connected(q int, p int) bool {
	return uf.arr[q] == uf.arr[p]
}

func (uf *QuickFInd) union(q int, p int) {
	var qVal = uf.arr[q]
	var pVal = uf.arr[p]

	for i := 0; i < len(uf.arr); i++ {
		if uf.arr[i] == pVal {
			uf.arr[i] = qVal
		}
	}
}

func main() {
	var uf = QuickFIndNew(5)
	uf.union(0, 2)
	uf.union(1, 2)
	fmt.Println(uf.connected(0, 2))
}
