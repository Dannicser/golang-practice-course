package main

import "fmt"

func main() {
	name := "Danil"
	pointer := &name // pointer не содержит значение, он содержит ссылку на место в памяти

	// anotherName := "Nataleck"

	fmt.Println("Memory location is", pointer)      // Memory location is 0x14eb6fa1a070
	fmt.Printf("Memory location is %p \n", pointer) // Memory location is 0x14eb6fa1a070
	fmt.Printf("Memory location is %p \n", pointer) // Memory location is 0x14eb6fa1a070

	fmt.Println("---------------")

	anotherName := name

	fmt.Println("Memory location is", &anotherName) // Memory location is 0x14eb6fa1a070
	fmt.Println("Memory location is", &anotherName) // Memory location is 0x14eb6fa1a070
	fmt.Println("Memory location is", &anotherName) // Memory location is 0x14eb6fa1a070

}
