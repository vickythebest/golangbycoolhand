package main

import "fmt"

func main() {
	arr := []int{1, 15, 3, 43, 8, 6, 7, 18, 9, 10}

	fmt.Print("New Array : ")
	// Display(arr)
	sortAscendingOrder(arr)
}

func sortAscendingOrder(arr []int) {

	length := len(arr) - 1
	left := 0
	right := 1

	for left < length {

		if arr[left] > arr[right] {
			temp := arr[right]
			arr[right] = arr[left]
			arr[left] = temp
			right++
		} else {
			right++
		}

		if left <= right && right == length {
			left += 1
			right = left + 1
		}
		fmt.Printf("Left :%v Right : %v \n", left, right)
		// Display(arr)
	}
	fmt.Print("Swape :")
	// Display(arr)
}
