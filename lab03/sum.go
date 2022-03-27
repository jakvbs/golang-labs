package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1 := 0
	num2 := 0

	if len(os.Args) == 1 {
		fmt.Print("Please enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Please enter second number: ")
		fmt.Scan(&num2)
	} else if len(os.Args) == 2 {
		num1, _ = strconv.Atoi(os.Args[1])
		fmt.Print("Please enter second number: ")
		fmt.Scan(&num2)
	} else {
		num1, _ = strconv.Atoi(os.Args[1])
		num2, _ = strconv.Atoi(os.Args[2])
	}

	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)

}
