package main

import "fmt"

func main() {
	var cupsQwy int = 3 // explicit type declaration
	var bottlesQwy = 3  // implicit type declaration

	// cupsQwy = "w" // compile-type error

	fmt.Println(cupsQwy, bottlesQwy)
}
