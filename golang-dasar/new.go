package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat *Address = new(Address)
	alamat2 := alamat

	alamat2.Country = "Indonesia"
	fmt.Println(alamat)
	fmt.Println(alamat2)
}