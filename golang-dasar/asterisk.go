package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"bogor", "jawa barat", "ID"}
	var address2 *Address = &address1 //pointer (menggunakan *)
	address2.City = "jakarta"
	fmt.Println(address1)
	fmt.Println(address2)

	address2 = &Address{"bandung", "jawa barat", "ID"}//tidak lgi reference ke address1
	fmt.Println(address1) //maka address 1 tetap jakarta
	fmt.Println(address2)

	address3 := &address1
	*address3 = Address{"depok", "jawa barat", "ID"} //jika menggunakan * di depan
	fmt.Println(address1) // maka address1 ikut berubah
	fmt.Println(address3)
}