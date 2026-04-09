package main

import "fmt"

func getUserInfo(name string, age int) {
	fmt.Printf("My name is %s. I am %d\n", name, age)
}

func main() {
	getUserInfo("Dan", 23)
	getUserInfo("Ivan", 22)
}
