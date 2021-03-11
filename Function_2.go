package main

import "fmt"

func test(myfunc func(string) string) {

	fmt.Println(myfunc("Kali Linux "))
}

func main() {

	test2 := func(name string) string {

		return "hello " + name
	}

	greeting := test2("John Wick")
	fmt.Println(greeting)

	test(test2)
}
