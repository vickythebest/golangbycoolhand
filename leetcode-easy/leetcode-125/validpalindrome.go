package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("abcba"))
}

func isPalindrome(s string) bool {

	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	str := regex.ReplaceAllString(s, "")

	str1 := strings.ToLower(str)

	left := 0
	right := len(str1) - 1
	for left < right {
		// fmt.Println(" str1[lenght] : ", str1[lenght])
		if string(str1[left]) != string(str1[right]) {
			return false
		}
		left++
		right--
	}

	// if len(str1)%2 == 0 {
	// 	fmt.Println("Even lenght of string : ", str1)

	// } else {
	// 	fmt.Println("Odd lenght of string : ", str1)
	// }
	return true
}
