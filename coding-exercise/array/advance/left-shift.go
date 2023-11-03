package main

import "fmt"

func LeftShift(array []int, index int) {
	length := len(array)
	index %= length
	fmt.Println("index : ", index)
	leftShift := make([]int, length)
	pos := 0
	for i := range array {
		if i >= index {
			leftShift[pos] = array[i]
			pos++
		}
	}
	display(leftShift)
}
