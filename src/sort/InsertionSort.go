package sort

import (
	"fmt"
)

// InsertionSort - typr for the sort
type InsertionSort struct {
	arr [5]int
}

func (insert *InsertionSort) swap(x, y int) {
	z := insert.arr[x]
	insert.arr[x] = insert.arr[y]
	insert.arr[y] = z
}

// InsertionSort - sorts array in ascending order.
func (insert *InsertionSort) InsertionSort() {
	for outer := 0; outer < len(insert.arr)-1; outer++ {
		for inner := outer + 1; inner > 0; inner-- {
			if insert.arr[inner] < insert.arr[inner-1] {
				insert.swap(inner-1, inner)
			} else {
				break
			}
		}
		fmt.Println(insert.arr)
	}
}
