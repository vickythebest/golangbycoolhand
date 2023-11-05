package main

/*
Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
	Input: nums = [4,3,2,7,8,2,3,1]
	Output: [5,6]
Example 2:
	Input: nums = [1,1]
	Output: [2]
*/

func multipleMissingElements(nums []int) []int {

	length := len(nums)
	var dummy []int
	var result []int

	for i := 1; i <= length; i++ {
		dummy = append(dummy, i)
	}
	var isExist = func(arr []int, num int) bool {
		for i := range arr {
			if arr[i] == num {

				return true
			}
		}
		return false
	}

	for _, num := range dummy {
		if !isExist(nums, num) {
			result = append(result, num)
		}
	}
	return result
}
