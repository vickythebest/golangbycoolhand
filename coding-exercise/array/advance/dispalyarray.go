package main

import (
	"fmt"
)

func display(array []int) {

	for i := range array {
		fmt.Print("\t", array[i])
	}
	fmt.Println()
}
