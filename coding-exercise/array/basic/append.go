package main

import "fmt"

func AppendElement(key int) {

	arr := []int{12, 4, 1, 5, 6, 8}

	arr = append(arr, key)
	for i := range arr {
		fmt.Println("Appended output :", arr[i])
	}

}
