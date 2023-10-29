package main

import "fmt"

//  data = [[4, 0, 5], [3, 5], [2, 4, 6], [8, 2, 9], [5, 6], [0, 4, 5], [4, 3], [17], [12, 1, 8]]

//  Print a list of the sums of the items in data.# [9, 8, 12, 19, 11, 9, 7, 17, 21]

func printSum(data [][]int) {

	list := []int{}
	for i := 0; i < len(data); i++ {
		sum := 0
		for j := 0; j < len(data[i]); j++ {

			sum = sum + data[i][j]
		}
		list = append(list, sum)
	}
	fmt.Print(" ", list)
}

func main() {

	data := [][]int{{4, 0, 5}, {3, 5}, {2, 4, 6}, {8, 2, 9}, {5, 6}, {0, 4, 5}, {4, 3}, {17}, {12, 1, 8}}

	printSum(data)
}
