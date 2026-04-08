package main

import "fmt"

func main() {
	var coffeName string = "Espresso"
	var coffeSize string = "Small"
	var coffePrice float32 = 2.15

	fmt.Printf("%s %s price is $%.3f \n", coffeSize, coffeName, coffePrice)
}
