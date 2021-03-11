package main

import "fmt"

func main() {

	// Immmutable ==> int , bool ,string , array
	var num1 int = 10
	var num2 int = num1
	num2 = 20

	fmt.Println(num1, num2)

	// mutabel slice , map

	var list []int = []int{1, 2, 3, 4}
	list2 := list
	list2[3] = 20
	fmt.Println(list, list2)

}
