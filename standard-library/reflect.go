package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` //structtag
}

type Person struct {
	Name string `required:"true"`
	Age  int `required:"true"`
	Email string `required:"true"`
}

func readFile(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name: ",valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println(field.Name, "with type", field.Type)
		fmt.Println(field.Tag.Get("required"))
		fmt.Println(field.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	// readFile(Sample{"Fahmi"})
	person := Person{
		Name:  "Fahmi",
		Age:   20,
		Email: "fahmi@go.dev",
	}

	fmt.Println(isValid(person))

}

//tag sangat cocok untuk validation sebuah data