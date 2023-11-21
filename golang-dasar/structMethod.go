package main

import "fmt"

type User struct {
	name string
	age  int
}

func (user User) sayHello(name string) {
	fmt.Println("hello", name, "my name is", user.name)
}

func main() {
	user := User{"fahmi", 26}
	User.sayHello(user, "budi")
}