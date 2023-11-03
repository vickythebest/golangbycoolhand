package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	array2 := []int{112, 21, 31, 43, 51, 61, 73, 81, 91, 101}

	fmt.Print("Array 1: ")
	display(array)
	fmt.Print("Array 2: ")
	display(array2)

	// fmt.Print("Shift element to left : ")
	// LeftShift(array, 15)

	// fmt.Print("Shift element to Right : ")
	// rightShift(array, 1)

	// fmt.Print("Reerse Array : ")
	// reverseOrswap(array)
	// reverseArray(array)

	// fmt.Print("Rotate Array : ")
	// rotateArray(array, 3)

	fmt.Print("Merge Array Result:")
	mergeArray(array, array2)

	fmt.Print("Merge sorted Array Result:")
	mergeSortedArray(array, array2)
}
