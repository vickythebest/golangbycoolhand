package main

import "fmt"

func (A1 *Array) insertElementInArray(key int) {

	A1.A = append(A1.A, key)

	fmt.Printf("New element inserted, \n display array :\n")
	A1.display()
}
