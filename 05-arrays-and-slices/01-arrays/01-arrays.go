package main

import "fmt"

func main() {
	var sizes [3]string // = [3]string{"", "", ""}

	fmt.Println(sizes) // [ ] - две пустых строки по умолчанию

	sizes[0] = "Small"
	sizes[1] = "Medium"
	sizes[2] = "Large"

	fmt.Println(sizes)
	fmt.Println(sizes[2])

	fmt.Println(3 != len(sizes))
}
