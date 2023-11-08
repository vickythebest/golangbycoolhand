package main

/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.
You must solve the problem without modifying the array nums and uses only constant extra space.

Example 1:

Input: nums = [1,3,4,2,2]
Output: 2
Example 2:

Input: nums = [3,1,3,4,2]
Output: 3
*/

func findDuplicate(nums []int) []int {

	length := len(nums)
	if length <= 0 {
		return []int{0}
	}

	result := make([]int, length)
	for i := 0; i < len(nums); i++ {
		result[i] = -1
	}

	index := 0
	result[index] = nums[0]
	index++

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result[index] = nums[i]
			index++
		}
	}
	return result
}
