package main

import "fmt"

func retunFunc(x string) func() {
	return func() {
		fmt.Println(x)

	}
}

func main() {

	retunFunc("Return a function")()
}
