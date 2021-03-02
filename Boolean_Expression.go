package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20

	val := num1 > num2
	// val := num1 >= num2
	// val := num1 < num2
	// val := num1 <= num2
	// val := num1 == num2
	// val := num1 != num2

	fmt.Printf("%t\n", val)

	// > , >= , < , <= ,== , !=

	state1 := true
	state2 := false

	state := state1 && state2
	// && => and
	// || => or
	// ! => not

	fmt.Printf("%t", state)

}
