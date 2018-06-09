package search

// BinarySearch - type to hold an array
type BinarySearch struct {
	arr [5]int
}

// BinarySearch - Finds an element in the sorted array
func (search *BinarySearch) BinarySearch(n int) (int, bool) {
	min := 0
	max := len(search.arr) - 1
	mid := min + (max-min)/2
	for min <= max {
		if search.arr[mid] == n {
			return mid, true
		} else if search.arr[mid] > n {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1, false
}
