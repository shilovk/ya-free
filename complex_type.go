package main

import (
	"fmt"
)

func main() {
	// Создание комплексного числа
	z := complex(3, 4) // 3 + 4i

	// Получение действительной и мнимой частей
	realPart := real(z)
	imaginaryPart := imag(z)

	fmt.Printf("Действительная часть: %f\n", realPart)
	fmt.Printf("Мнимая часть: %f\n", imaginaryPart)

	// Арифметические операции с комплексными числами
	anotherComplex := complex(1, 2)
	sum := z + anotherComplex
	product := z * anotherComplex

	fmt.Printf("Сумма: %v\n", sum)
	fmt.Printf("Произведение: %v\n", product)
}
