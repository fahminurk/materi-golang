package main

import "fmt"

func main() {

	name := "budi"

	if name == "fahmi" {
		fmt.Println("hello", name)
	} else {
		fmt.Println("you're not fahmi")
	}

	

	if length := len(name); length > 10 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("terlalu pendek")
	}
}