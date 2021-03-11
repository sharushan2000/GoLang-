package main

import "fmt"

func changeValue(state *bool) {
	*state = false
}

func main() {

	/*num1 := 22
	num2 := &num1
	fmt.Println(num1, *num2)
	fmt.Println(&num1, num2)*/

	x := true
	fmt.Println(x)
	changeValue(&x)
	fmt.Println(x)

}
