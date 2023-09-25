package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sqrt(x int) int {
	fmt.Println("Sqrt of ", x)
	right := int(math.Sqrt(float64(math.MaxInt)))
	// fmt.Println(right)
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
		// fmt.Println(middle)

		result := middle * middle
		// fmt.Println("middle ", middle, " result ", result)
		if result == float32(x) {
			return int(middle)
		} else if result < float32(x) {
			left = int(middle) + 1
			// fmt.Println("Left : ", left)
		} else {
			right = int(middle) - 1
			// fmt.Println("Right : ", right)
		}

	}
	return right
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter positive number : ")
	scanner.Scan()
	digit, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("The sqrt of ", scanner.Text(), "= ", sqrt(int(digit)))
	fmt.Println("Result : ", sqrt(8))
}
