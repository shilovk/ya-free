package main

import "fmt"

func increment(x int) {
	// x — локальная переменная для этой функции
	x++
}

func Divide(x int) (half int) {
	half = x / 2
	return // тогда в инструкции return имя можно не указывать
}

func Sum(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}

func Index(st string, a rune) (index int, ok bool) {
	for i, c := range st {
		if c == a {
			return i, true
		}
	}
	return // вернутся значения по умолчанию
}

func fact(n int) int {
	if n == 0 { // терминальная ветка — то есть условие выхода из рекурсии
		return 1
	} else { // рекурсивная ветка
		return n * fact(n-1)
	}
}

func Fib(n int) int {
	switch {
	case n <= 1: // терминальная ветка
		return n
	default: // рекурсивная ветка
		return Fib(n-1) + Fib(n-2)
	}
}

func factIter(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}
func FibIter(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return a
}

func main() {
	n := 5
	// n копируется в переменную x
	increment(n) // значение n не изменится
	fmt.Println(n)

	f := Divide(20)
	fmt.Println(f)

	fmt.Println(Sum(1, 2))

	fmt.Println(Index("Строка", rune('а')))

	fmt.Println(fact(5))
	fmt.Println(factIter(5))

	fmt.Println(Fib(14))
	fmt.Println(FibIter(14))
}
