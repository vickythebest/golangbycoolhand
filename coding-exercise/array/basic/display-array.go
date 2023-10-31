package main

import "fmt"

func (A1 *Array) display() {

	for i := range A1.A {

		fmt.Println("Display Element : ", A1.A[i])
	}
	fmt.Println("Done !")
}
