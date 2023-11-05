package advance

import "fmt"

func rightShift(array []int, positions int) {

	length := len(array)
	if positions == 0 || len(array) == 0 {
		fmt.Println("Invalid array or positions")
		return
	}

	rightShift := make([]int, length)

	positions %= length

	fmt.Println("positions : ", positions)

	for i := range array {
		if i < (length - positions) {
			rightShift[i] = array[i]
		}
	}

	display(rightShift)
}
