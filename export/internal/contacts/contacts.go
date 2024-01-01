package contacts

import "fmt"

const Email = "support@example.com" // глобальная экспортируемая константа

var support string // глобальная неэкспортируемая переменная

// SetSupport устанавливает значение переменной support.
func SetSupport(s string) { // экспортируемая функция
	support = s
}

// GetContact возвращает имя и email.
func GetContact() string { // экспортируемая функция
	return fmt.Sprintf("%s <%s>", support, Email)
}
