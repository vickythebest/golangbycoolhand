package main

func findMinAndMax(array []int) (int, int) {

	min := array[0]
	max := array[0]
	i := 0

	for i < len(array) {
		if min > array[i] {
			min = array[i]
			i++
		} else if max < array[i] {
			max = array[i]
			i++
		} else {
			i++
		}
	}
	return min, max
}
