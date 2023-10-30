package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {

	if n == 0 {
		return 1 // return 1 in case of multiplicaion otherwise 0 in addition or subtraction
	}
	return n * factorial(n-1)

}
