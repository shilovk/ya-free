package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]string) // создаём map — о применении функции make для создания переменных типа map будет рассказано ниже
	m["foo"] = "bar"             // заполняем парами «ключ-значение»
	m["ping"] = "pong"
	fmt.Println(m) // печатаем

	type MyMap map[string]string

	var m1 MyMap
	m1 = make(MyMap, 5)

	// объект готов
	m1["foo"] = "bar"
	fmt.Println(m1)

	MyStringMap := map[string]string{"first": "первый", "second": "второй"}
	fmt.Println(MyStringMap)

	var mmm map[string]string
	if mmm != nil { // если не проверить это условие,
		mmm["foo"] = "bar" // то здесь можно получить panic
	}
	fmt.Println("mmm", mmm) // mmm map[]

	m2 := map[int]string{1: "first"}
	v, ok := m2[1]
	fmt.Println(v, ok)
	delete(m2, 2) // ошибки не будет
	delete(m2, 1)
	v, ok = m2[1]
	fmt.Println(v, ok)

	m3 := make(map[string]string)
	m3["foo"] = "bar"
	m3["bazz"] = "yup"
	for k, v := range m3 {
		// k будет перебирать ключи,
		// v — соответствующие этим ключам значения
		fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
	}
	for k, _ := range m3 {
		m3[k] = "here key " + k // применяем к таблице индексное выражение
		// и модифицируем её прямым доступом к ячейкам
	}
	fmt.Println(m3)
	for k, _ := range m3 { // обратите внимание на подчёркивание _
		delete(m3, k)
	}
	fmt.Println(m3)

	// предложение
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}

	//Теперь попробуйте сами. Сделайте map для хранения прейскуранта в рублях:
	//хлеб — 50,
	//молоко — 100,
	//масло — 200,
	//колбаса — 500,
	//соль — 20,
	//огурцы — 200,
	//сыр — 600,
	//ветчина — 700,
	//буженина — 900,
	//помидоры — 250,
	//рыба — 300,
	//хамон — 1500.
	//Выведите перечень деликатесов — продуктов дороже 500 рублей.
	//Заказ выдан слайсом []string{"хлеб", "буженина", "сыр", "огурцы"}.
	//Посчитайте стоимость заказа.
	strProducts := `
		хлеб — 50,
		молоко — 100,
		масло — 200,
		колбаса — 500,
		соль — 20,
		огурцы — 200,
		сыр — 600,
		ветчина — 700,
		буженина — 900,
		помидоры — 250,
		рыба — 300,
		хамон — 1500.
	`
	// Задаем регулярное выражение для поиска пробелов
	regex := regexp.MustCompile(`(\s+|\.)`)
	// Удаляем пробелы из строки с использованием регулярного выражения
	strProducts = regex.ReplaceAllString(strProducts, "")
	partsProducts := strings.Split(strProducts, ",")
	products := make(map[string]int)
	for _, v := range partsProducts {
		split := strings.Split(v, "—")

		// Преобразование строки в int
		price, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Ошибка при преобразовании:", err)
			return
		}

		products[split[0]] = price
	}
	fmt.Println(products)

	richProducts := make([]string, 0, len(products))
	for k, v := range products {
		if v > 500 {
			richProducts = append(richProducts, k)
		}
	}
	fmt.Println(richProducts)

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	price := 0
	for _, v := range order {
		price += products[v]
	}
	fmt.Println(price)
}
