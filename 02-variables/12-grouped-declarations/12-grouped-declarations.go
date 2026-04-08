package main

import "fmt"

func main() {
	const d = 1 // ошибка компиляции не вознимает, потому что const

	var (
		a int    = 1
		b string = ""
		c bool   = true
	)

	fmt.Println(a, b, c)
}
