package main

import (
	"fmt"
)

type MergeSort struct {
}

func (merge *MergeSort) split(arr []int, first []int, last []int) {
	for index := 0; index < len(arr); index++ {
		if index < len(first) {
			first[index] = arr[index]
		} else {
			last[index-len(first)] = arr[index]
		}
	}
}

func (merge *MergeSort) merge(first []int, last []int, arr []int) {
	firstIndex := 0
	lastIndex := 0
	arrIndex := 0

	for firstIndex < len(first) && lastIndex < len(last) {
		if first[firstIndex] <= last[lastIndex] {
			arr[arrIndex] = first[firstIndex]
			firstIndex++
		} else {
			arr[arrIndex] = last[lastIndex]
			lastIndex++
		}
		arrIndex++
	}
	if firstIndex < len(first) {
		for firstIndex < len(first) {
			arr[arrIndex] = first[firstIndex]
			arrIndex++
			firstIndex++
		}
	} else {
		for lastIndex < len(last) {
			arr[arrIndex] = last[lastIndex]
			arrIndex++
			lastIndex++
		}
	}
}

// MergeSort - Sorts a array using merge sort algo.
func (merge *MergeSort) MergeSort(arr []int) {
	if len(arr) == 1 {
		return
	}

	midIndex := len(arr)/2 + len(arr)%2
	first := make([]int, midIndex)
	last := make([]int, len(arr)-midIndex)
	merge.split(arr, first, last)

	merge.MergeSort(first)
	merge.MergeSort(last)

	fmt.Println(first)
	fmt.Println(last)

	merge.merge(first, last, arr)

	fmt.Println(arr)
}

func main() {
	merge := MergeSort{}
	arr := [5]int{5, 4, 3, 2, 1}
	x := arr[0:5]
	merge.MergeSort(x)
}
