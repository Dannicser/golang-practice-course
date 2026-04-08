package main

import "fmt"

func main() {
	// untyped constant with intenger value
	const a = 10

	// untyped constant with float value
	const b = 2.1

	fmt.Printf("type is %T %T\n", a, b)

	const c = a + b // ошибки нет, несмотря на то что мы добавляем int + float, если явно не типизировать

	fmt.Println(c)
}
