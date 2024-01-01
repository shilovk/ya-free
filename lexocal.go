package main

import "fmt"

func outer() {
	// Область видимости переменной a - только внутри функции outer
	var a int = 5

	// Внутри функции inner можно обратиться к переменной a
	inner := func() {
		fmt.Println(a)
	}

	inner()
}

func main() {
	outer()
}
