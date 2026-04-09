package main

import "fmt"

func main() {
	var name, age = getUserInfo()

	fmt.Println(name, age)
}

func getUserInfo() (string, int) {
	return "Dan", 23
}
