package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("New Array : ")
	Display(array)

	// fmt.Print("Shift element to left : ")
	// LeftShift(array, 15)

	// fmt.Print("Shift element to Right : ")
	// rightShift(array, 1)

	fmt.Print("Reerse Array : ")
	reverseOrswap(array)
	reverseArray(array)

	fmt.Print("Rotate Array : ")
	rotateArray(array, 3)
}
