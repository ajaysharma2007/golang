package sort

import (
	"fmt"
)

// BubbleSort - Sorts the list is ascending order
type BubbleSort struct {
	arr [5]int
}

func (bubble *BubbleSort) swap(x, y int) {
	z := bubble.arr[x]
	bubble.arr[x] = bubble.arr[y]
	bubble.arr[y] = z
}

// BubbleSort - Sorts the array in ascendign order.
func (bubble *BubbleSort) BubbleSort() {
	for outer := 0; outer < len(bubble.arr)-1; outer++ {
		swapped := false
		for inner := 0; inner < len(bubble.arr)-outer-1; inner++ {
			if bubble.arr[inner] > bubble.arr[inner+1] {
				bubble.swap(inner, inner+1)
				swapped = true
			}
		}
		fmt.Println(bubble.arr)
		if !swapped {
			break
		}
	}
}
