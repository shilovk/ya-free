package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func Generate(seed int) func() {
	return func() {
		fmt.Println(seed) // замыкание получает внешнюю переменную seed
		seed += 2         // переменная модифицируется
	}

}

func fib() func() int {
	x1, x2 := 0, 1
	// возвращаемая функция замыкает x1, x2
	return func() int {
		x1, x2 = x2, x1+x2
		return x1
	}
}

// countCall — функция-обёртка для подсчёта вызовов
func countCall(f func(string)) func(string) {
	cnt := 0
	// получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		cnt++
		fmt.Printf("Функция %s вызвана %d раз\n", funcname, cnt)
		f(s)
	}
}

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now() // получаем текущее время
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start)) // получаем интервал времени как разницу между двумя временными метками
	}
}

func myprint(s string) {
	fmt.Println(s)
}

func main() {
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	iterator()
	iterator()

	f := fib()       // Получили функцию-замыкание. f() — захватила x1, x2. x1 = 0, x2 = 1
	fmt.Println(f()) // x1 = 1, x2 = 1
	fmt.Println(f()) // x1 = 1, x2 = 2
	fmt.Println(f()) // x1 = 2, x2 = 3
	fmt.Println(f()) // x1 = 3, x2 = 5
	fmt.Println(f()) // x1 = 5, x2 = 8
	fmt.Println(f()) // x1 = 8, x2 = 13

	// Создадим две функции-обёртки, одна из которых будет подсчитывать количество вызовов,
	// а вторая — время исполнения функции.
	countedPrint := countCall(myprint)
	countedPrint("Hello world")
	countedPrint("Hi")
	// Обратите внимание, что мы оборачиваем уже обёрнутую функцию,
	//а значение счётчика вызовов при этом сохраняется
	countAndMetricPrint := metricTimeCall(countedPrint)
	countAndMetricPrint("Hello")
	countAndMetricPrint("World")
	//Функция main.myprint вызвана 1 раз
	//Hello world
	//Функция main.myprint вызвана 2 раз
	//Hi
	//Функция main.myprint вызвана 3 раз
	//Hello
	//Execution time:  3.833µs
	//Функция main.myprint вызвана 4 раз
	//World
	//Execution time:  3.583µs
}
