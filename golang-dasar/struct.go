//struct sama seperti interface / kumpulan type data

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	isActive      bool
}

func main() {
	var customer Customer
	customer.Name = "fahmi"
	customer.Address = "jalan raya"
	customer.Age = 21
	customer.isActive = true
	fmt.Println(customer)

	fahmi := Customer{
		Name:    "fahmi",
		Address: "jalan raya",
		Age:     21,
		isActive: true,
	}
	fmt.Println(fahmi)
}