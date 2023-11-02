package main

import "fmt"

func (a *Array) binarySearch(key int, start int, end int) {

	if start < end {
		mid := (start + end) / 2
		// fmt.Printf("start from index : %v and end index : %v ", start, end)

		if key == a.A[mid] {
			fmt.Printf("Index of element %v at index : %v: ", key, mid)
			return
		} else if key < a.A[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		a.binarySearch(key, start, end)
	}
}
