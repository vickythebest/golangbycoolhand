package main

import "fmt"

func (a *Array) deleteElement(key int) {

	var array = Array{}

	for i := range a.A {

		if a.A[i] != key {

			array.A = append(array.A, a.A[i])
		}
	}
	fmt.Println("Elements has been deleted, here is the new Array : ")
	array.display()
}
