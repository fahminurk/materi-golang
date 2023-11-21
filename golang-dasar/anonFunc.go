package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {

	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
blackList := func(name string) bool {
	return name == "anjing"
}

registerUser("fahmi", blackList)
registerUser("anjing", blackList)
}