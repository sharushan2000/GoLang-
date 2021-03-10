package main

import "fmt"

func main() {
	num := 10

	switch {
	case num > 10:
		fmt.Println("num is greater than ten")

	case num < 10:
		fmt.Println("num is less than 10")

	default:
		fmt.Println("num is equal to ten")

	}

}
