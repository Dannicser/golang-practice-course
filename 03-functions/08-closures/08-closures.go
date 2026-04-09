package main

import "fmt"

var createTempAdjuster = func(baseTemp int) (int, func(change int) int) {
	var adjustTemp = func(change int) int {
		baseTemp = baseTemp + change

		return baseTemp

	}

	return baseTemp, adjustTemp
}

func main() {
	var baseTemp, adjuster = createTempAdjuster(10)

	fmt.Println(adjuster(-1))
	fmt.Println(adjuster(-1))
	fmt.Println(adjuster(-1))
	fmt.Println(adjuster(-1))

	fmt.Println(baseTemp) // не изменится, так как мы присвоили его выше и больше не трогали !!!

}
