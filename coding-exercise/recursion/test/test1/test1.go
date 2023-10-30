package main

import "fmt"

func main() {
	a := 2048
	sum := 0

	foo(a, sum)

	fmt.Printf("%d\n", sum)
}

func foo(n int, sum int) {

	k := 0
	j := 0

	if n == 0 {
		return
	}

	k = n % 10
	j = n / 10

	sum = sum + k

	foo(j, sum)

	fmt.Printf("“%d”", k)
}
