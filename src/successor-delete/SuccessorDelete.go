package main

import (
	"fmt"
	ufm "union-find-max"
)

// SuccessorDelete type
type SuccessorDelete struct {
	uf        *ufm.UnionFindMax
	size      int
	statusArr []bool
}

// NewSuccessorDelete Instance
func NewSuccessorDelete(n int) *SuccessorDelete {
	newStatusArr := make([]bool, n, n)
	return &SuccessorDelete{ufm.NewUnionFindMax(n), n, newStatusArr}
}

func (sd *SuccessorDelete) remove(n int) int {
	sd.statusArr[n] = true
	if n == sd.size-1 {
		return n
	}
	return sd.uf.Find(n)
}

func (sd *SuccessorDelete) successor(n int) int {
	var index int
	index = sd.uf.Find(n)

	if sd.statusArr[index] == true {
		index = -1
	}

	return index
}

func main() {
	var sd = NewSuccessorDelete(10)

	sd.remove(2)
	fmt.Println(sd.successor(2))

	sd.remove(3)
	fmt.Println(sd.successor(2))

	fmt.Println(sd.successor(8))
	sd.remove(8)

	fmt.Println(sd.successor(8))

	sd.remove(9)
	fmt.Println(sd.successor(8))
	fmt.Println(sd.successor(9))

	sd.remove(5)
	sd.remove(4)

	fmt.Println(sd.successor(3))

}
