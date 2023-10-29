package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.



Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter first string : ")
	scanner.Scan()
	input1 := scanner.Text()

	fmt.Print("Please enter second string : ")
	scanner.Scan()
	input2 := scanner.Text()

	// fmt.Println(lengthOfLongestSubstring(input1))
	// longestCommonSubsequence(input1, input2)

	fmt.Println("Result : ", longestCommonSubsequence(input1, input2))

	fmt.Print("Result recursion: ")
	fmt.Print(lcs(input1, input2, len(input1), len(input2)))

}
func longestCommonSubsequence(text1 string, text2 string) int {

	rows := len(text1)
	col := len(text2)

	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	// fmt.Printf("Text1 size = %v and Text2 size = %v", rows, col)
	fmt.Println()

	var max = func(x, y int) int {
		// fmt.Println("check max between ", x, " and ", y)
		if x < y {
			return y
		}
		return x
	}
	// Fill the DP table using the update rule
	for i := 1; i <= rows; i++ {
		for j := 1; j <= col; j++ {
			// fmt.Println("text1[i-1] :", text1[i-1], " text2[j-1] :", text2[j-1])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// fmt.Println("retun result : ", dp[rows][col])
	return dp[rows][col]
}

// RECURSION
func lcs(text1, text2 string, m, n int) int {

	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	if m == 0 || n == 0 {
		return 0
	}

	if text1[m-1] == text2[n-1] {
		return 1 + lcs(text1, text2, m-1, n-1)
	}

	return max(lcs(text1, text2, m, n-1), lcs(text1, text2, m-1, n))
}
