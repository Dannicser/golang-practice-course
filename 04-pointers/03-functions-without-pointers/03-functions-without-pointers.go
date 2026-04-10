package main

func main() {
	println(calcDiscount(50))
	println(calcDiscount(100))
	println(calcDiscount(200))
}

func calcDiscount(price int) (res int) {
	const discountRate = 0.13

	res = price - int(float64(price)*discountRate)

	return

}
