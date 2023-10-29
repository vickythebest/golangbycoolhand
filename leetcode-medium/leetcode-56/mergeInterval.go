package main

import (
	"fmt"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print(" input please : ")
	// scanner.Scan()

	// userinput := scanner.Text()
	// fmt.Print(" input please : ", userinput)
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	merge(input)

}

func merge(intervals [][]int) [][]int {
	fmt.Println(" merge")

	for i := 0; i < len(intervals); i++ {
		fmt.Println("I : ", i)
		// {{1, 3}, {2, 6}, {8, 10}, {15, 18}}

		for j := 0; j < len(intervals[i]); j++ {
			fmt.Println(" J : ", intervals[i][j])
		}
	}

	fmt.Print("for done")
	output := [][]int{{1, 6}, {2, 6}, {8, 10}, {15, 18}}
	return output
}
