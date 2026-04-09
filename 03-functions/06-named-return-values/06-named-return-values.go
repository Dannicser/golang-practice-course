package main

import "fmt"

// можно объявить возвращаемое значение

func calc(a int, b int) (c int) {
	c = a + b

	// naked return
	return // тогда оно автоматически вернется
}

func main() {
	fmt.Println(calc(2, 3))
}
