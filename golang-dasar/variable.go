package main

import "fmt"

func main() {
	var name string
	var age = 26
	babu := "udin petot" // hanya dapat digunakan sat deklarasi pertamakali
	
	name = "fahmi nurkamil"
	fmt.Println(name)

	name = "asd"
	fmt.Println(name)

	fmt.Println(age)
	fmt.Println(babu)

//multiple variable var/const
	var(
		firstName = "fahmi"
		lastName = "nurkamil"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	const pi = 3.14

	
}

//setiap variable jika telah di deklar 
// harus dipakai kecuali menggunakan const, kalau tidak error
