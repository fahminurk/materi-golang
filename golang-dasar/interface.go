//kontrak

package main

import "fmt"

type Person struct {
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

func main(){
	person := Person{Name: "fahmi"}
	sayHello(person)
	fmt.Println(person.GetName())
}

func sayHello(hasName HasName) {
	fmt.Println("Hello, " + hasName.GetName())
}

type HasName interface {
	GetName() string
}