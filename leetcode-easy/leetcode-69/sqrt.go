package main

import (
	"fmt"
	"math"
)

func sqrt(x int) int {
	fmt.Println("Get Sqrt of ", x)
	right := int(math.Sqrt(float64(math.MaxInt)))
	fmt.Println(right)
	var left = 0

	for left <= +right {

		if left*left == x {
			fmt.Println("sqrt : ", left)
			return left
		}

		if right*right == x {
			return right
		}

		middle := float32(left + ((right - left) / 2))
		fmt.Println(middle)

		result := middle * middle
		// fmt.Println("middle ", middle, " result ", result)
		if result == float32(x) {
			return int(middle)
		} else if result < float32(x) {
			left = int(middle) + 1
			fmt.Println("Left : ", left)
		} else {
			right = int(middle) - 1
			fmt.Println("Right : ", right)
		}

	}
	return right
}

func main() {
	fmt.Println(sqrt(8))
}
