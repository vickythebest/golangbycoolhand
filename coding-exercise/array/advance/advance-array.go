package main

import (
	"fmt"
)

func main() {

	array1 := []int{1, 21, 3, 4, 51, 6, 7, 81, 9, 10}

	array2 := []int{11, 21, 31, 43, 51, 61, 75, 71, 91, 101}

	fmt.Print("Array 1: ")
	display(array1)
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

	// fmt.Print("Merge Array Result:")
	// mergeArray(array1, array2)

	// fmt.Print("Merge sorted Array Result:")
	// mergeSortedArray(array1, array2)

	fmt.Print("Union Array Result:")
	unionArray(array1, array2)

	// union(array1, array2)

}
