package main

import "fmt"

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

func main() {
	myFigure := square
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	// Вычисление площади для конкретной фигуры
	myArea := ar(5.0)
	fmt.Printf("Площадь фигуры: %f\n", myArea)
}
