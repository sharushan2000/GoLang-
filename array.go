package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr)

	fmt.Println("--------------------------------------------------------")

	arr_2 := [5]int{1, 2, 3, 4, 5}
	for _, num := range arr_2 {
		fmt.Println(num)
	}

	fmt.Println("--------------------------------------------------------")

	Array_2D := [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	for _, d1 := range Array_2D {
		for _, d2 := range d1 {
			fmt.Println(d2)
		}
	}

}
