package main

func (a *Array) findMaxElement() int {

	max := a.A[0]

	for i := range a.A {
		if a.A[i] > max {
			max = a.A[i]
		}
	}
	return max
}
