package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//...
	var a string
	a = "абц"
	println(utf8.RuneCountInString(a)) // 3

	// Использование руны для представления символа 'A'
	var myRune rune = 'A'

	// Использование руны в строке
	myString := "🕰️Привет, мир!"
	firstChar := rune(myString[0])
	//str := strconv.Itoa(int(firstChar))
	//str := fmt.Sprintf("%c", firstChar)

	fmt.Printf("Руна: %c\n", myRune)
	fmt.Printf("Первый символ строки: %c\n", firstChar)
	fmt.Printf("Первый символ строки: \n", myString)
}
