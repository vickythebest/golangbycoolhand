package main

import "fmt"

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Constraints:

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109
*/

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3

	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	length := -1

	k := ((m + n) - 1)

	left := -1
	right := -1

	m = m - 1
	n = n - 1
	for length <= k {

		if m >= 0 {
			left = nums1[m]
		}

		if n >= 0 {
			right = nums2[n]
		}

		fmt.Printf("left : %v right : %v K : %v \n", left, right, k)

		if left >= right {
			nums1[k] = left
			m--
			k--

		} else {
			nums1[k] = right
			n--
			k--

		}
		length++
		fmt.Printf("M : %v  N : %v K : %v length : %v \n", m, n, k, length)

	}

	for _, result := range nums1 {
		fmt.Println("oupput : ", result)
	}
}

// func merge(nums1 []int, m int, nums2 []int, n int) {

// 	i := 0
// 	j := 0
// 	// newarray := num1
// 	var mergeArray []int
// 	length := 0
// 	for len(mergeArray) <= (m+n)-1 {

// 		fmt.Printf("num1[%v] == num2[%v] and length : %v", nums1[i], nums2[j], length)
// 		fmt.Println()

// 		if i < m && nums1[i] < nums2[j] {
// 			fmt.Printf("num1[%v] ", nums1[i])
// 			mergeArray = append(mergeArray, nums1[i])
// 			i++
// 			fmt.Printf(" I : %v", i)
// 			fmt.Println()
// 		} else {
// 			fmt.Printf("num2[%v] ", nums2[j])
// 			mergeArray = append(mergeArray, nums2[j])
// 			j++
// 			fmt.Printf(" J : %v", j)
// 			fmt.Println()
// 		}
// 		length++
// 	}

// 	// nums1 = mergeArray
// 	for i, result := range mergeArray {
// 		nums1[i] = result

// 		fmt.Println("output : ", result)
// 	}
// }

// func merge(nums1 []int, m int, nums2 []int, n int) {

// 	var length int = (m + n)
// 	// fmt.Println("length : ", length, "  len(nums1) : ", len(nums1))

// 	var mergeArray []int

// 	// fmt.Println("mergeArryay length : ", len(mergeArray))
// 	for i := 0; i <= length-1; i++ {
// 		// fmt.Println("I : ", i, nums1[i])
// 		for j := i; j <= n-1; j++ {
// 			fmt.Println(" J : ", j, nums2[j])
// 			if nums2[j] <= nums1[i] {
// 				mergeArray = append(mergeArray, nums2[j])
// 			}
// 		}
// 		mergeArray = append(mergeArray, nums1[i])
// 		fmt.Println("----->", mergeArray[i])
// 	}

// 	// for _, result := range mergeArray {
// 	// 	fmt.Println(" result : ", result)
// 	// }
// }
