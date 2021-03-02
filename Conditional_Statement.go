package main

import "fmt"

func main() {
	var num int

	num = 10

	if num > 100 {
		fmt.Println("Greater than 100 ")

	} else if num > 50 {
		fmt.Println("Greater than 50 ")

	} else {
		fmt.Println("Less than 50 ")
	}

}
