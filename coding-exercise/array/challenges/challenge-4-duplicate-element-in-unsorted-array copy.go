package main

import "fmt"

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
// Solution 0
func findDuplicateNumber(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		fmt.Printf("Slow : %v , Fast : %v ", slow, fast)
		fmt.Println()

		if slow == fast {
			break
		}

	}

	fmt.Println("Phase 2: Move the slow pointer to the entrance of the loop")
	fmt.Println("Before : ", slow)
	slow = nums[0]
	fmt.Println("After : ", slow)
	for slow != fast {
		fmt.Printf("Slow : %v , Fast : %v ", slow, fast)
		fmt.Println()
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// Solution 1
func findDuplicateNumber1(nums []int) int {

	for i := range nums {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

// Solution 2
func findDuplicateNumber2(nums []int) []int {

	// length := len(nums)
	// create new array with same length
	result := []int{}

	var max = func(arr []int) int {

		max := 0
		for i := range arr {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max + 1
	}
	// initiate array element to 0
	for i := 0; i < max(nums); i++ {
		result = append(result, 0)
	}
	// fmt.Println(result)

	// mark elements -1 if it already scanned
	count := 0
	for num := 0; num < len(nums); num++ {

		if result[nums[num]] == 0 {
			result[nums[num]] = -1
		} else {
			count++
			result[nums[num]] = count
		}
	}
	for i := 1; i < len(result); i++ {
		if result[i] > 0 {
			fmt.Println("Duplicate number is : ", i)
		}
	}
	return result
}
