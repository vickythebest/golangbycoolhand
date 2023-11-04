package main

import "fmt"

func interSection(arr1 []int, arr2 []int) {
	length := len(arr1)
	k := 0
	interSection := make([]int, length)
	// fmt.Println(interSection[k])
	for i := range arr1 {
		for j := range arr2 {
			if arr1[i] == arr2[j] {
				interSection[k] = arr1[i]
				k++
			}
		}
	}
	fmt.Println("Result")
	display(interSection)
	// i := 0
	// j := 0
	// for i < len(arr1) && j < len(arr2) {

	// 	if arr1[i] == arr2[j] {
	// 		interSection[k] = arr1[i]
	// 	}else{
	// 		i++
	// 	}
	// }
}
