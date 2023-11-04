package main

/*
	array := []int{1, 21, 3, 4, 51, 6, 7, 81, 9, 10}

	array2 := []int{11, 21, 31, 43, 51, 61, 73, 81, 91, 101}
*/
func unionArray(arr1 []int, arr2 []int) {

	length := len(arr1) + len(arr2)

	u := 0
	unionArray := make([]int, length)
	for i := range arr1 {
		unionArray[u] = arr1[i]
		u++
	}
	var isNotContain = func(arr []int, key int) bool {
		for i := range arr {
			if arr[i] == key {
				return false
			}
		}
		return true
	}

	for j := range arr2 {

		if isNotContain(unionArray, arr2[j]) {
			unionArray[u] = arr2[j]
			u++
		}
	}

	display(unionArray)
}
