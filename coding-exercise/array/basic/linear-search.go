package main

func (a *Array) searchElement(key int) int {

	for i := range a.A {

		if a.A[i] == key {
			return i
		}
	}
	return -1
}
