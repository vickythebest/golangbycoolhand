package main

/*
Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8

Input: nums = [3,0,1]
Output: 2

*/
func singleMissingElement(arr []int) int {

	length := len(arr)

	var addElements = func(arr []int) int {
		sum := 0
		for i := range arr {
			sum = sum + arr[i]
		}
		return sum
	}

	total := addElements(arr)

	// fmt.Println("Max number in array ", max)
	sum := sum(length)

	// fmt.Println("Sum of elements in array ", sum)

	return sum - total
}

func sum(max int) int {
	if max == 1 {
		return max
	} else if max == 0 {
		return max
	} else {
		return max + sum(max-1)
	}
}
