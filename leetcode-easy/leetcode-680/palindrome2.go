package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string : ")
	scanner.Scan()
	fmt.Println("user input : ", scanner.Text(), " is :", validpalindrome(scanner.Text()))

	// fmt.Println(validpalindrome("abcba"))
	// fmt.Println(validpalindrome(string()))
}

func validpalindrome(str string) bool {

	isValid := func(str string, left int, right int) bool {
		// left := 0
		// right := (len(str) - 1)
		// fmt.Printf(" substring : %v left : %v < right : %v <> ", str, left, right)
		fmt.Println()
		for left < right {
			// fmt.Println(" isValid <> : False")
			if string(str[left]) != string(str[right]) {
				// fmt.Println(" isValid : False")
				return false
			}
			left++
			right--
		}
		return true
	}
	left := 0
	right := (len(str) - 1)

	for left < right {
		if string(str[left]) != string(str[right]) {
			// return isValid(str[left:right-1]) || isValid(str[left+1:right])
			return isValid(str, left+1, right) || isValid(str, left, right-1)
		}
		left++
		right--
	}
	return true
}
