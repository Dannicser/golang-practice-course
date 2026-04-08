package main

import "fmt"

func main() {
	var name = "Danil"
	var age = 23
	var isFrontender = true
	var account = 2.4

	fmt.Printf("type name is %s \n", name)
	fmt.Printf("type age is %T \n", age)
	fmt.Printf("type isFrontender is %T \n", isFrontender)
	fmt.Printf("type account is %T \n", account)
}
