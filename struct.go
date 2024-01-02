package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"` // - означает, что это поле не будет сериализовано
}

type Tree struct {
	Value      int
	LeftChild  *Tree
	RightChild *Tree
}

func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
	return Person{
		Name:        name,
		Email:       email,
		DateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}

type GetUserRequest struct {
	UserId    string `json:"user_id" yaml:"user_id" format:"uuid" example:"2e263a90-b74b-11eb-8529-0242ac130003"`
	IsDeleted *bool  `json:"is_deleted,omitempty" yaml:"is_deleted"`
}

func main() {
	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	p := Person{
		Name:        "MyName",
		Email:       "mail@mail.com",
		DateOfBirth: date,
	}
	// или
	var t Tree
	man := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}

	fmt.Println("Person:", p)
	fmt.Println("NewPerson:", man)
	fmt.Printf("JsonPerson: %v", string(jsMan))
	fmt.Println("Tree:", t)
	fmt.Printf("Tree: %#v", t) //Tree: main.Tree{Value:0, LeftChild:(*main.Tree)(nil), RightChild:(*main.Tree)(nil)}%

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}

	reqRaw, _ := json.Marshal(req)
	os.Stdout.WriteString("\n")
	os.Stdout.Write(reqRaw)

	var c1 struct{}
	// или
	c2 := struct{}{}
	fmt.Println(unsafe.Sizeof(c1))   //0
	fmt.Println(unsafe.Pointer(&c1)) //0x102a90d20
	fmt.Println(unsafe.Sizeof(c2))   //0
	fmt.Println(unsafe.Pointer(&c2)) //0x102a90d20
}
