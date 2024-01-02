package main

import (
	"fmt"
	"ya-free/export/internal/foo"
	// импортируем пакет из поддиректории internal/contacts
	"ya-free/export/internal/contacts"
)

func main() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)

	// f := foo.privateFoo{} // ошибка компиляции
	f := foo.NewPrivateFoo()
	fmt.Println(f.Value) // поле Value экспортируемое, то есть его можно использовать
}
