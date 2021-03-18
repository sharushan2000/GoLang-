package main

import "fmt"

type intro interface {
	greeting() (string, int)
}

type father struct {
	name string
	age  int
}

type mother struct {
	name string
	age  int
}

func (f father) greeting() (string, int) {
	greeting := "I'm " + f.name + " and I'm "

	return greeting, f.age
}

func (m mother) greeting() (string, int) {
	greeting := "I'm " + m.name + " and I'm "

	return greeting, m.age
}

func main() {

	f := father{"John Wick", 47}
	m := mother{"Hermione", 28}

	fmt.Println(f.greeting())
	fmt.Println(m.greeting())

	introductions := []intro{f, m}

	for _, introduction := range introductions {
		fmt.Println(introduction.greeting())
	}

}
