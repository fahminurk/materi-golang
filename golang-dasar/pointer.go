package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta"
	fmt.Println(address1) //pas by value, seperti shallow copy di js address1 tidak berubah
	fmt.Println(address2)

	address3 := &address1 //refference menggunakan & (pointer)
	address3.City = "Bandung"

	fmt.Println(address1) //address1 ikut berubah
	fmt.Println(address3)
}
