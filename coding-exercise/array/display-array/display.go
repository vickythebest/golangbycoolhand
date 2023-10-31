package main

import "fmt"

type Array struct {
	A      []int
	Size   int
	Length int
}

func main() {

	array := &Array{
		A:      []int{10, 20, 30},
		Size:   10,
		Length: 10,
	}
	array.display()
}

func (A1 *Array) display() {
	for i := range A1.A {
		fmt.Println("Elements : ", i)
	}
}
