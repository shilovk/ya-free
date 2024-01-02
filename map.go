package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func uniqueValues(input []int) []int {
	uniqueMap := make(map[int]bool)
	uniqueSlice := []int{}

	for _, element := range input {
		if _, exists := uniqueMap[element]; !exists {
			uniqueMap[element] = true
			uniqueSlice = append(uniqueSlice, element)
		}
	}

	return uniqueSlice
}

func findMy(arr []int, k int) map[int]int {
	result := make(map[int]int)
	for _, v1 := range arr {
		for _, v2 := range arr {
			if ((v1 + v2) == k) && (result[v2] == 0) {
				result[v1] = v2
			}
		}
	}
	return result
}

// ya-variant
func findYa(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func removeDuplicates(input []string) []string {
	m := make(map[string]int)
	for _, v := range input {
		m[v] += 1
	}
	s := make([]string, 0, len(m))
	for v, _ := range m {
		s = append(s, v)
	}
	return s
}

func RemoveDuplicatesYa(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)

		}
		inputSet[v] = struct{}{}
	}

	return output
}

// как можно заметить, алгоритм пройдётся по массиву всего один раз
// если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений

// как можно заметить, алгоритм пройдётся по массиву всего один раз
// если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений

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

	//Ассоциативные массивы широко применяются при решении алгоритмических задач.
	//Когда количество данных более нескольких десятков,
	//поиск значения в map происходит эффективнее, чем в массиве.
	//Опираясь на эту информацию, попробуйте решить следующую задачу,
	//которую часто предлагают на собеседованиях.
	//Дан массив целых чисел A и целое число k.
	//Нужно найти и вывести индексы пары чисел, сумма которых равна k.
	//Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке.
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 7
	fmt.Println("findMy", findMy(arr, k))
	fmt.Println("findYa", findMy(arr, k))

	result := []int{}
	for _, v1 := range arr {
		for _, v2 := range arr {
			if (v1 != v2) && ((v1 + v2) == k) {
				result = append(result, v1, v2)
			}
		}
	}
	fmt.Println(uniqueValues(result))

	//Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println("source", input)
	fmt.Println("removeDuplicates", removeDuplicates(input))
}
