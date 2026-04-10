package main

import "fmt"

func main() {
	nums := [4]int{1, 2, 3, 4}

	sliceOne := nums[0:2]  // первое значение включительно, второе невключительно [1,2)
	sliceTwo := nums[:]    // [:] весь диапазон
	sliceThree := nums[1:] // до конца

	fmt.Println(sliceOne, sliceTwo, sliceThree)
}
