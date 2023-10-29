package main

/*
Given a string s, find the longest palindromic subsequence's length in s.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements
without changing the order of the remaining elements.
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string to check plandrom")
	scanner.Scan()

	input := scanner.Text()

	longestPalindromeSubseq(input)
}

func longestPalindromeSubseq(s string) int {

	// left := 0
	right := len(s)
	count := 0

	dp := make([][]int, right)

	for i := range dp {
		dp[i] = make([]int, right)
	}

	var max = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	for i := 0; i < right-1; i++ {
		for j := right - 1; j > 0; j-- {
			fmt.Println("start : ", string(s[i]), " end : ", string(s[j]))
			if s[i] == s[j] {
				fmt.Println("match 1: ", string(s[i]), " hatch 2 : ", string(s[j]))
				dp[i][j] = dp[i+1][j-1] + 2
				count++
			} else {
				// fmt.Printf("I : %v, J : %v : ", i, j)
				// fmt.Println("dp[i+1][j] : ", dp[i+1][j])

				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	fmt.Print("Length : ", dp[0][right-1])
	return dp[0][right-1]
}
