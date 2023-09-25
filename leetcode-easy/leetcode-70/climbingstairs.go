package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter stair number : ")
	scanner.Scan()
	count := scanner.Text()
	stair, _ := strconv.ParseInt(count, 10, 64)
	fmt.Println(int(stair))
	fmt.Println("Total distinct way to climb stair : ", claimbstair(int(stair)))
	// fmt.Println("Total distinct way to climb stair : ", claimbstair(5))
}

func claimbstair(num int) int {

	var one, two = 1, 1

	for num != 0 {
		temp := one
		one = one + two
		two = temp
		num -= 1
	}

	return two
}

// func claimbstair(num int) int {

// 	count := 0

// 	if num > 0 {
// 		count = count + 1
// 	}
// 	reminder := num % 2
// 	result := (num / 2)

// 	// fmt.Println("Reminder : ", reminder, " result : ", result)

// 	if reminder == 0 && (result*2) == num {
// 		count = count + 1
// 	}

// 	if reminder == 1 {
// 		result1 := result
// 		reminder1 := reminder
// 		for reminder1 < result1 {
// 			reminder1 = reminder1 + 2
// 			if (reminder1 + result1) == num {
// 				count = count + 1
// 			}
// 		}
// 	}

// 	for result >= 1 {
// 		reminder = reminder + 2
// 		if (reminder + result) == num {
// 			count = count + 1
// 		}
// 		result = result - 1
// 	}

// 	return count
// }
