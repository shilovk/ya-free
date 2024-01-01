package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(day Weekday) Weekday {
	return (day % 7) + 1
}

func main() {
	type Name string
	type Fruit string

	var fruit Fruit
	var name Name

	fruit = "Apple"
	// Чтобы привести один тип к другому,
	// в Go используется такой синтаксис: type(variable). Проиллюстрируем на предыдущем примере:
	name = Name(fruit) // так, после приведения типов, работает
	println(name)

	var str string
	str = "Hello, world!"
	println(string(str[0])) //72
	// приведения типов:
	//int64Var := int64(5)
	//float32Var := float32(101.3)

	// типы rune и byte представляют собой алиасы к int32 и uint8:}
	// type rune = int32
	// type byte = uint8
	type MyString = string // MyString здесь — это псевдоним типа string

	var a string
	var b MyString
	a = b // ошибки нет
	println(a)

	// определите переменные ver, id, pi
	ver := "v0.0.1"
	var id int64 = 0
	var pi = 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
	//ver = v0.0.1 id = 0 pi = 3.1415

	const (
		Black = iota
		Gray
		White
	)

	// счётчик обнуляется
	const (
		Yellow = iota
		Red
		Green = iota // это присваивание не обнулит iota
		Blue
	)

	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue)

	const (
		_ = iota * 10 // обратите внимание, что можно пропускать константы
		ten
		twenty
		thirty
	)

	const (
		hello = "Hello, world!" // iota равна 0
		one   = 1               // iota равна 1

		black = iota // iota равна 2
		gray
	)

	fmt.Println(ten, twenty, thirty)
	fmt.Println(black, gray)

	var today Weekday = Sunday
	tomorrow := NextDay(today)
	fmt.Println("today =", today, "tomorrow =", tomorrow)

	// логическое НЕ
	// возвращается одна переменная типа bool
	//a := false
	//if !a {}
	//
	//// логическое И
	//var a, b int
	//if a == 1 && b == 2 {}
	//
	//// исключающее ИЛИ (XOR)
	//var a, b bool
	//if (a || b) && !(a && b) {}

	aa, bb := 1, 0

	incB := func() bool {
		bb = bb + 1
		return true
	}

	if aa == 1 || incB() {
		fmt.Println("Hello")
	}

	fmt.Println(aa, bb) // 1, 0

	aaa := 0.10000001 // float64
	// инициализация и основное условие
	if bbb := float32(aaa); bbb > float32(0.1) {
		fmt.Println(bbb)
		fmt.Println("Var a is GT float32(0.1)")
	}

	dd := -100
	switch {
	case dd > 0:
		if dd%2 == -0 {
			break
		}
		fmt.Println("Odd positive value received")
	case dd < 0:
		fmt.Println("Negative value received")
		fallthrough
	default:
		fmt.Println("Default value handling")
	}
}
