package main

import (
	"fmt"
	"time"
)

type UserPerson struct {
	Name        string
	Age         int
	lastVisited time.Time
}

func UpdatePersonWithLastVisited(d *UserPerson) {
	d.lastVisited = time.Now()
}

func main() {
	// UserPerson в другом пакете
	d := UserPerson{
		Name: "Alex",
		Age:  25,
	}
	UpdatePersonWithLastVisited(&d)
	fmt.Println(d)

	var a int = 5
	//var p *int
	p := &a
	fmt.Println(a, p) //a=5 p=0xc0000b2008

	//Литералы композитных типов создают в памяти переменную соответствующего типа,
	//поэтому указатель можно создать вот так:
	type A struct {
		IntField int
	}
	// Литерал А{} создаёт в памяти переменную типа А. Затем от неё берётся указатель
	z := &A{
		IntField: 10,
	}
	println(z)

	//А ещё в Go есть встроенная функция new().
	//В качестве параметра ей передаётся тип,
	//а возвращается указатель на новую переменную соответствующего типа.
	type T struct {
		IntField int
	}

	//Тип указателя на указатель описывается как **T, например **int.
	//Чтобы получить или изменить значение,
	//хранящееся по указателю, применяют оператор разыменования (dereference) *.
	ii := 42
	pp := &ii
	fmt.Println(pp)  // адрес
	fmt.Println(*pp) // читаем значение переменной ii через указатель p
	*pp = 21         // записываем в переменную ii значение 21 через указатель p
	fmt.Println(ii)  // 21

	incrementCopy := func(i int) {
		i++
	}

	increment := func(i *int) {
		*i++
	}

	i := 42

	incrementCopy(i)
	fmt.Println(i) // 42

	increment(&i)
	fmt.Println(i) // 43

	g := 1
	k := &g
	l := &k

	*k = 3
	**l = 5
	println(*k)  //5
	println(**l) //5
	println(g)   //5
	g += 4 + *k + **l

	fmt.Printf("%d", *k) //19
}
