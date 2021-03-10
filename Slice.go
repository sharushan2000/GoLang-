package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T \n", arr)

	println("-------------------------------------------")

	var s []int = arr[1:3]

	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Println(len(s))
	fmt.Printf("%T \n", s)

	println("-------------------------------------------")

	slice := []int{11, 22, 33, 44, 55}

	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	fmt.Printf("%T \n", slice)

	println("---------------------------------------------")

	slice = append(slice, 66)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	fmt.Printf("%T \n", slice)
}
