package main

import "fmt"

type Array struct {
	A      []int
	size   int
	length int
}

func main() {

	array := Array{
		A:      make([]int, 10),
		size:   10,
		length: 10,
	}

	fmt.Printf("Enter number of elements\n")
	n := 0
	fmt.Scanf("%d", &n)

	fmt.Println("N : ", n)
	if n > len(array.A) {
		fmt.Println("Input count exceeds array size.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scanf("\n%d", &array.A[i])
	}

	fmt.Println("Elements : ", array.A[:n])

}
