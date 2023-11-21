package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountry(address Address) {
	address.Country = "Indonesia"
}

func changeCountryPointer (address *Address){
	address.Country = "Indonesia"
}

func main() {
	alamat := Address{}
	changeCountry(alamat)
	fmt.Println(alamat) // tidak berubah karena by pas value, bukan reference

	changeCountryPointer(&alamat) //kasih tanda &
	fmt.Println(alamat)
}