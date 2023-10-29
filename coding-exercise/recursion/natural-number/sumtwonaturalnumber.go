package main

import "fmt"

func main() {

	fmt.Println(sumTwoNaturalNumber(5))
}

func sumTwoNaturalNumber(n int) int {

	if n < 0 {
		return 0
	}

	return sumTwoNaturalNumber(n-1) + n
}
