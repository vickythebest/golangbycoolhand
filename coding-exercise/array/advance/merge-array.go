package main

/*
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	array2 := []int{11, 21, 31, 43, 51, 61, 73, 81, 91, 101}
*/
func mergeArray(arr1 []int, arr2 []int) {

	length := len(arr1) + len(arr2)
	mergeArray := make([]int, length)
	// mergeArray := []int{}

	i := 0
	j := 0
	m := 0

	for i < len(arr1)-1 || j < len(arr2)-1 {

		if i < len(arr1) {
			mergeArray[m] = arr1[i]
			i++

		} else {
			mergeArray[m] = arr2[j]
			j++
		}
		m++
	}

	if i < len(arr1) {
		mergeArray[m] = arr1[i]
		i++
		m++
	}

	if j < len(arr2) {
		mergeArray[m] = arr2[j]
		j++
		m++
	}

	display(mergeArray)
}

func mergeSortedArray(arr1 []int, arr2 []int) {

	length := len(arr1) + len(arr2)
	mergeArray := make([]int, length)

	i := 0
	j := 0
	m := 0

	for i < len(arr1)-1 || j < len(arr2)-1 {

		if i < len(arr1) && arr1[i] < arr2[j] {
			mergeArray[m] = arr1[i]
			i++

		} else {
			mergeArray[m] = arr2[j]
			j++
		}
		m++
	}

	if i < len(arr1) {
		mergeArray[m] = arr1[i]
		i++
		m++
	}

	if j < len(arr2) {
		mergeArray[m] = arr2[j]
		j++
		m++
	}

	display(mergeArray)
}
