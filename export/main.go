package main

import (
	"fmt"
	// импортируем пакет из поддиректории internal/contacts
	"ya-free/export/internal/contacts"
)

func main() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
}
