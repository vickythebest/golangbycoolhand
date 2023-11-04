package main

import (
	"fmt"
)

func display(array []int) {

	for i := range array {
		if array[i] != 0 {
			fmt.Print("\t", array[i])
		}
	}
	fmt.Println()
}
