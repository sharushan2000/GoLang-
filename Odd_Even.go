package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Printf("Enter The Number (More than 0 ): ")
	fmt.Scanf("%d", &number)
	if number != 0 {
		if number%2 == 0 {
			fmt.Println("Even Number")
		} else {
			fmt.Println("Odd Number")
		}
	} else {
		fmt.Println("Enter a Valid Number  ")
	}

}
