package main

import "fmt"

//membuat alias/type untuk func
// type Filter func(string) string

func getGoodBye(name string) string {
	return "Good bye b" + name
}

func sayHelloWithFilter(name string, filter func(string) string) { //(name string, filter Filter)
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func SpamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}



func main() {
//func as variable
	bye := getGoodBye
	fmt.Println(bye("fahmi"))

	sayHelloWithFilter("anjing", SpamFilter)
	sayHelloWithFilter("budi", SpamFilter)
	
}