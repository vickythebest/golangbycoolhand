package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" Please enter digit : ")
	scanner.Scan()

	userinput, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("fibonacci of : %v is %v", int(userinput), fibonacci(int(userinput)))
	// num := 5
	// fmt.Printf("fibonacci of : %v is %v", num, fibonacci(num))

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
