package main

import "fmt"

// Базовая структура
type Animal struct {
	Name string
}

// Структура, встраивающая Animal
type Dog struct {
	Animal // Встраиваем Animal в Dog
	Breed  string
}

func main() {
	// Создаем экземпляр Dog
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// Мы можем обращаться как к полям Dog, так и к полям Animal
	fmt.Println("Name:", myDog.Name)
	fmt.Println("Breed:", myDog.Breed)
}
