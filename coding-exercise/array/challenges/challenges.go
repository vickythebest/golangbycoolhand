package main

import "fmt"

func main() {

	unSortedArray := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// array := []int{1, 1}
	sortedArray := []int{1, 2, 3, 4, 6, 8, 12}
	// sortedArray := []int{6, 7, 8, 9, 11, 12, 15, 16, 17, 20}

	fmt.Println(singleMissingElement(unSortedArray))

	multipleMissingElements(unSortedArray)

	findMissingElements(sortedArray)
}
