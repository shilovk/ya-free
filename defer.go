package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func EvaluationOrder() {
	defer fmt.Println("deferred")
	fmt.Println("evaluated")
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
	// Какие-то долгие вычисления
}

func PanicingFunc() {
	defer func() {
		if r := recover(); r != nil { // встроенная функция recover останавливает панику и возвращает описание произошедшего
			fmt.Println("Panic is caught", r)
		}
	}()
	///
	///

	panic("Мне здесь совсем ничего не нравится!")
	// встроенная функция panic () вызывает панику у функции.
	// в качестве аргумента ей принято передавать причину паники. Именно она будет возвращена функцией recover

}

func main() {
	EvaluationOrder()
	//evaluated
	//deferred

	fmt.Println("Hello")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
	//Hello
	//World
	//3
	//2
	//1

	a := "some text"
	defer func(s string) {
		fmt.Println(s)
	}(a)
	a = "another text"
	// some text

	// открываем файл
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрыть файл
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// работаем с файлом
	_, err = file.WriteString("Text")
	if err != nil {
		log.Fatal(err)
	}

	VeryLongTimeFunction()

	PanicingFunc()
}
