package main

import "fmt"

func (A1 *Array) replaceElement(in, out int) {

	for i := range A1.A {
		if A1.A[i] == out {
			A1.A[i] = in
		}
	}
	fmt.Printf(" %v element has been replace, \n display array :\n", out)
	A1.display()
}
