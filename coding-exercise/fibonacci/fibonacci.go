package main

import "fmt"

func main() {
	num := 5
	fmt.Printf("fibonacci of : %v is %v", num, fibonacci(num))

}

func fibonacci(num int) int {

	fmt.Println("Num : ", num)
	if num == 0 {
		fmt.Println("if num = ", num)
		return num
	} else if num == 1 {
		fmt.Println("else if num = ", num)
		return num
	} else {
		fmt.Println("else  num = ", num)
		return fibonacci(num-1) + fibonacci(num-2)
	}

}
