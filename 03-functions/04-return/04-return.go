package main

import "fmt"

func calcRewards(amountSpent int) int {
	return amountSpent * 2
}

func main() {
	var totalPoints int = 20

	var result int = totalPoints + calcRewards(20)

	fmt.Println(result)
}
