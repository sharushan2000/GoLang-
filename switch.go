package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Enter The Your Age : ")
	fmt.Scanf("%d", &age)

	switch age {
	case 0:
		fmt.Println("Enter a valid age  ")

	case 13:
		fmt.Println("Welcome to teenage ")

	case 21:
		fmt.Println("You are a adult now ")

	default:
		fmt.Println("Welcome ")

	}
}
