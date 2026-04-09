package main

import "fmt"

func main() {
	var calcTax = func(amount float64) float64 {
		const taxRate float64 = 0.10

		return amount * taxRate

	}

	const salary = 10000.0
	var net = salary - calcTax(salary)

	fmt.Println(net)
}
