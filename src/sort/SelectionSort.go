package sort

// SelectionSort - sorts a given array in ascending order
type SelectionSort struct {
	arr [3]int
}

func (selection *SelectionSort) swapIndex(x, y int) {
	z := selection.arr[x]
	selection.arr[x] = selection.arr[y]
	selection.arr[y] = z
}

// SelectionSort - sorts an given array using selection sort.
func (selection *SelectionSort) SelectionSort() {

	for outer := 0; outer < len(selection.arr)-1; outer++ {
		minIndex := outer
		for index := outer + 1; index < len(selection.arr); index++ {
			if selection.arr[minIndex] > selection.arr[index] {
				minIndex = index
			}
		}
		if minIndex != outer {
			selection.swapIndex(minIndex, outer)
		}
	}
}
