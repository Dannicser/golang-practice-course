package main

import "fmt"

func main() {
	var userTypes = [3]string{"Admin", "Viewer", "Guest"} // array literal

	fmt.Println(userTypes)
	fmt.Println(len(userTypes))
	fmt.Println(userTypes[len(userTypes)-1]) // получение последнего элемента в массиве
}
