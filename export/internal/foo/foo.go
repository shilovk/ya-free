package foo

// privateFoo — неэкспортируемый тип
type privateFoo struct {
	Value string
}

// NewPrivateFoo — конструктор типа privateFoo
// Функция публичная, то есть может быть вызвана из других пакетов
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data"}
}
