package main

import "fmt"

func test() {
	fmt.Println("Hello World!")
}

func math(x, y int) (int, int) {

	return x + y, x - y

}

func greeting(name string, year int) (age int) {
	defer fmt.Println("Bye")
	age = 2021 - year
	fmt.Println("Hello", name)
	return
}

func main() {
	test()

	add, sub := math(20, 10)
	fmt.Println(add, sub)

	age := greeting("sharu", 2000)
	fmt.Println(age)

}
