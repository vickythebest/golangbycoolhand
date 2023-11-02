package main

import (
	"fmt"
)

func Display(array []int) {

	for i := range array {
		fmt.Print("\t", array[i])
	}
	fmt.Println()
}
