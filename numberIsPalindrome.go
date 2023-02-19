package main

import (
	"fmt"
)

func main() {
	var number, remainder, reverse int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	temp := number

	// Reversing the number
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a palindrome\n", temp)
	} else {
		fmt.Printf("%d is not a palindrome\n", temp)
	}
}
