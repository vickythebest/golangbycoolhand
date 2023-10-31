package main

import (
	"fmt"
	"strings"
)

type Array struct {
	A      []int
	size   int
	length int
}

func main() {

	fmt.Printf("Enter number of elements\n")
	n := 0
	fmt.Scanf("%d", &n)

	array := Array{
		A:      make([]int, n),
		size:   10,
		length: 2,
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scanf("\n%d", &array.A[i])
	}

	fmt.Println()
	if n >= len(array.A) {
		fmt.Println("Input count exceeds array size.")
	}

	fmt.Println(" *********** Display Array  *********** ")
	array.display()

	var userChose string
	var newElement int
	fmt.Print("Do you wants to insert more element (Yes / No) : ")

	choice := false
	fmt.Scan(&userChose)

	if strings.ToLower(userChose) == "yes" {
		choice = true
		fmt.Printf("Enter new elements\n")
		fmt.Scan(&newElement)
	}

	if choice {
		fmt.Printf("Insert %v into array ", newElement)
		array.insertElementInArray(newElement)
		// AppendElement(newElement)
	}

	fmt.Printf("Do you want to replace element (Yes / No):")
	fmt.Scan(&userChose)

	if strings.ToLower(userChose) == "yes" {

		fmt.Printf("Enter new element :")
		fmt.Scan(&n)

		var out int
		fmt.Printf("Enter existing element which you wants to replace :")
		fmt.Scan(&out)
		array.replaceElement(n, out)

	}

	fmt.Printf("Enter element which you want to delete : ")
	fmt.Scan(&n)
	array.deleteElement(n)

}
