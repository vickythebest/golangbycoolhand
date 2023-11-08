package main

import "fmt"

func main() {

	// unSortedArray := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// array := []int{1, 1}
	// sortedArray := []int{1, 2, 3, 4, 6, 8, 12}
	// sortedArray := []int{6, 7, 8, 9, 11, 12, 15, 16, 17, 20}

	// duplicateNumber := []int{1, 3, 4, 2, 2}
	// duplicateNumber := []int{2, 2, 2}
	// duplicateNumber := []int{2, 1, 1, 1, 4}

	// fmt.Println(singleMissingElement(unSortedArray))

	// multipleMissingElements(unSortedArray)

	// findMissingElements(sortedArray)

	// result := findDuplicateNumber(duplicateNumber)

	// fmt.Println(result)

	// duplicate := []int{1, 1, 2, 2, 3, 4, 4}
	// fmt.Println(duplicate)
	// fmt.Println()
	// result := findDuplicate(duplicate)
	// fmt.Println()
	// fmt.Println("Result : ", result)

	minMax := []int{-5, -2, -6, -9, -1, -3}
	min, max := findMinAndMax(minMax)
	fmt.Printf("Min : %v ,Max : %v ", min, max)

}
