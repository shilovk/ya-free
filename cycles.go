package main

import "fmt"

func main() {
	// создаём переменную
	v := 0
	for i := 1; i < 10; i++ {
		// наращиваем переменную
		v++
	}
	fmt.Println(v)

	for a, b := 5, 10; a < 10 && b < 20; a, b = a+1, b+2 {
		// do stuff
	}

	// Цикл с одним условием
	//начальная инициализация
	i := 0
	for i < 5 {
		// выводим результат на экран
		fmt.Println(i)
		// наращиваем переменную
		i++
	}

	// бесконечный циклы
	//for {{
	//for ;; {}
	//for ; true; {}

	// Цикл range
	// создаём массив
	array := [3]int{11, 12, 13}
	// итерируемся
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}

	sum, limit := 0, 100
	for i := 0; true; i++ {
		if i%2 != 0 {
			continue // переход к следующему числу, так как i — нечётное
		}

		if sum+i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

		sum += i
	}
	fmt.Println(sum)

	//Метки
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
	fmt.Println("End")
	//[0, 0]
	//End

outerLoopLabelСontinue:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabelСontinue
		}
	}
	fmt.Println("End")
	//  [0, 0]
	//	[1, 0]
	//	[2, 0]
	//	[3, 0]
	//	[4, 0]
	//	End

	//вывести чётные числа в диапазоне [0:20] с указанием десятка.
	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break // break относится к ближайшему switch
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}
	//01: 2
	//01: 4
	//01: 6
	//01: 8
	//02: 12
	//02: 14
	//02: 16
	//02: 18
}
