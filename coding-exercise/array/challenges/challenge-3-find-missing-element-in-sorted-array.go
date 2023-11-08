package main

import "fmt"

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
	Input: nums = [1, 2, 3, 4, 6, 8, 12]
	Output: [5,7,9,10,11]
*/

func findMissingElements(nums []int) {

	result := []int{}
	// length := len(nums)
	diff := (nums[0] - 0)
	// fmt.Println("diff : ", diff)
	for i, num := range nums {
		// fmt.Printf("Num : %v , Index : %v ", num, i)
		newDiff := (num - i)
		if newDiff != diff {
			// fmt.Println("New Diff : ", newDiff)
			for diff < newDiff {
				result = append(result, diff+i)
				diff++
			}
		}
	}
	fmt.Println("Missing Element : ", result)
}
