package main

import "fmt"

func reverseOrswap(array []int) {

	i := 0
	j := len(array) - 1

	for i < j {

		temp := array[i]
		array[i] = array[j]
		array[j] = temp
		i++
		j--

	}
	Display(array)
}

func reverseArray(array []int) {
	fmt.Print("Before :")
	Display(array)
	length := len(array)

	for i := 0; i < length/2; i++ {
		array[i], array[length-i-1] = array[length-i-1], array[i]

	}
	fmt.Print("After :")
	Display(array)
}
