package sort

import (
	"fmt"
)

// ShellSort - Sorts a given array
type ShellSort struct {
	arr [5]int
}

func (shell *ShellSort) swapIndex(x, y int) {
	z := shell.arr[x]
	shell.arr[x] = shell.arr[y]
	shell.arr[y] = z
}

func (shell *ShellSort) ShellSort() {
	for increment := len(shell.arr) / 2; increment >= 1; increment /= 2 {
		fmt.Println("Increment is : ", increment)
		for startIndex := 0; startIndex < increment; startIndex++ {
			for outer := startIndex; outer < len(shell.arr)-increment; outer += increment {
				for inner := outer + increment; inner-increment >= 0; inner -= increment {
					if shell.arr[inner] < shell.arr[inner-increment] {
						shell.swapIndex(inner, inner-increment)
					} else {
						break
					}
				}
			}
			fmt.Println(shell.arr)
		}
	}
}
