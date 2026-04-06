package main

import (
	"fmt"

	"github.com/fatih/color" // будем обращаться к пакету, используя последнюю часть
)

// чтобы код запустился, обязательно должна быть функция main, это точка входа
func main() {
	color.Yellow("Hello World!")
	fmt.Println("Hello World!")
}
