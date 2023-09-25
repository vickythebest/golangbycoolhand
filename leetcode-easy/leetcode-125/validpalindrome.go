package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("enter string to validate palindrome : ")
	scanner.Scan()
	fmt.Println("The string palindrom validation : ", isPalindrome(scanner.Text()))

	// fmt.Println(isPalindrome("abcba"))
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
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
