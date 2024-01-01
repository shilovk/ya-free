package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func reverseSlice(input []int) []int {
	length := len(input)
	reversed := make([]int, length)
	copy(reversed, input)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Сумма чисел:", result)

	//mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
	mySlice1 := make([]int, 0)     // слайс [], базовый массив []
	mySlice2 := make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	mySlice3 := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]
	mySlice4 := []int{1, 2, 3}     // [1 2 3]
	fmt.Println(mySlice1, mySlice2, mySlice3, mySlice4)

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]
	weekTempArr[0] = 555

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7

	sl := []int{1, 2, 3} // [1 2 3]
	sl = sl[:len(sl)-1]  // [1 2] Ёмкость массива при этом не изменяется
	fmt.Println(sl)

	a := []int{1, 2, 3, 4}
	b := a[2:3] // b = [3]
	b = append(b, 7)
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
	b = append(b, 8, 9, 10)
	b[0] = 11
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6

	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
	// 1. Создаём слайс s с базовым массивом на 7 элементов.
	// Четыре первых элемента будут доступны в слайсе.

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
	// 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
	// Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7),
	// то базовый массив остаётся прежним.
	// Слайс s тоже изменился, но его длина осталась прежней

	slice2 := append(s[1:2], 7)
	fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
	// 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

	slice3 := append(s, slice1[1:]...)
	fmt.Println(len(slice3), cap(slice3)) // 8 14
	// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,
	// что больше ёмкости базового массива.
	// Будет создан новый базовый массив ёмкостью 14,
	// ёмкость нового базового массива подбирается автоматически
	// и зависит от текущего размера и количества добавленных элементов

	// 5. Легко проверить, что slice3 ссылается на новый базовый массив
	s[1] = 99
	fmt.Println(s, slice1, slice2, slice3)
	// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]

	s1 := []int{5, 4, 1, 3, 2}
	sort.Ints(s1)
	fmt.Println(s1) // [1 2 3 4 5]

	bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
	fmt.Printf("%s", bytes.TrimSpace(bSlice)) // a lone gopher
	fmt.Printf("%s", bSlice)                  // \t\n a lone gopher \n\t\r\n

	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]

	s111 := []int{1, 2, 3}
	if len(s111) != 0 { // защищаемся от паники
		s111 = s111[:len(s)-1]
	}
	fmt.Println(s111) // [1 2]

	s222 := []int{1, 2, 3}
	if len(s222) != 0 { // защищаемся от паники
		s222 = s222[1:]
	}
	fmt.Println(s222) // [2 3]

	s333 := []int{1, 2, 3, 4, 5}
	i333 := 2
	if len(s333) != 0 && i333 < len(s333) { // защищаемся от паники
		s333 = append(s333[:i333], s[i333+1:]...)
	}
	fmt.Println(s333) // [1 2 4 5]

	s_1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s_1, s2)) // false
	fmt.Println(reflect.DeepEqual(s_1, s3)) // false
	fmt.Println(reflect.DeepEqual(s_1, s4)) // true

	// Создаем массив
	array := [5]int{1, 2, 3, 4, 5}

	// Создаем срез на основе массива
	slice := array[1:4]

	// Выводим срез
	fmt.Println("Массив:", array)
	fmt.Println("Срез:", slice)

	//Создайте слайс и заполните его числами от 1 до 100.
	//Оставьте в слайсе первые и последние 10 элементов и
	//разверните слайс в обратном порядке.
	//Подумайте, можно ли подобным образом развернуть строку?
	dim := 100
	ssss := make([]int, 0, dim)
	// заполняем слайс числами
	for i := 0; i < dim; i++ {
		ssss = append(ssss, i+1)
	}
	// оставляем первые и последние 10 элементов
	ssss = append(ssss[:10], ssss[dim-10:]...)
	dim = len(ssss)
	// разворачиваем слайс
	for i := range ssss[:dim/2] {
		ssss[i], ssss[dim-i-1] = ssss[dim-i-1], ssss[i]
	}
	fmt.Println(ssss)
}
