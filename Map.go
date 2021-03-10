package main

import "fmt"

func main() {

	var info map[string]int

	info = map[string]int{"sharu": 20, "kumar": 29, "ajith": 33}

	fmt.Println(info)

	//chage value

	info["sharu"] = 100

	fmt.Println(info)

	//add new key

	info["vj"] = 200
	fmt.Println(info)

	//delele

	delete(info, "ajith")

	// info := make(map[string]int)
}
