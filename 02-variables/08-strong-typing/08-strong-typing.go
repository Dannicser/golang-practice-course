package main

import "fmt"

func main() {
	price := 100.0 // нельзя 100, это разные типы, автоматической конвертации не будет
	qwt := 4.2

	fmt.Println("Total is", price*qwt)
}
