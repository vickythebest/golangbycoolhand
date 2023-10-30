package main

import "fmt"

func main() {
	fmt.Println(powerOf(2, 5))
}

func powerOf(n int, m int) int {
	if m == 0 {
		return 1
	}

	return n * powerOf(n, m-1)
}
