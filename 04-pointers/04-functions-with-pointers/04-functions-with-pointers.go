package main

import "fmt"

func main() {
	var price = 100
	const discountRate = 0.13

	fmt.Println("Место в памяти аргумента функции - ", &price) // Разные места в памяти

	price = calcDiscount(price, discountRate) // Две разные переменные, просто значение скопированное

	fmt.Println(price)

	fmt.Println("----------")

	fmt.Println("Price before applyDiscount - ", price)

	applyDiscount(&price) // прокидываем указатель, а не переменную

	fmt.Println("Price after applyDiscount using pointers - ", price)
}

// Две разные переменные, просто значение скопированное
func calcDiscount(price int, discountRate float64) int {
	fmt.Println("Место в памяти параметра функции - ", &price) // Разные места в памяти, просто значение скопированное

	return price - int(float64(price)*discountRate)
}

func applyDiscount(price *int) {
	*price = 1000
}
