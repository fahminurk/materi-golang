package main

import "fmt"

func main() {
	// cara pertama
	var person map[string]string = map[string]string{}
	person["name"] = "fahmi"
	person["age"] = "26"

	// cara cepat
	data := map[string]string{
		"name": "fahmi",
		"age":  "26",
	}

	fmt.Println(person)
	fmt.Println(data)
	fmt.Println(data["name"])
	fmt.Println(data["tidakadakey"]) //string kosong tergantung type data yg di declare

	// function map
	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "fahmi"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}

// map tidak ada batasan pada length