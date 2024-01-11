package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	FName  string
	LName  string
	City   string
	Mobile int64
}

func structIterator() {
	fields := reflect.VisibleFields(reflect.TypeOf(Student{}))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}
}

//Key: FName      Type: string
//Key: LName      Type: string
//Key: City       Type: string
//Key: Mobile     Type: int64

func main() {
	s := Student{"Chan", "Tulsa", "Bangalore", 7777777777}
	//s := struct{Foo string; Bar int }{"foo", 2}

	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
	//Field: FName    Value: Chan
	//Field: LName    Value: Tulsa
	//Field: City     Value: Bangalore
	//Field: Mobile   Value: 7777777777

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
	//[Chan Tulsa Bangalore 7777777777]

	structIterator()
}
