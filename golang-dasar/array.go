package main

import "fmt"

func main() {
	var names [3]string //array dengan jumlah data 3
	names[0] = "fahmi"
	names[1] = "nurkamil"
	names[2] = "udin"
	// names[3] = "petot" //error melebihi batas
	fmt.Println(names)

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)

	var values2 = [...]int{ //jumlahnya/lengthnya tidak pasti tapi harus langsung di deklar isinya
		1,
		2,
		3,
		4,
	}

	fmt.Println(len(values2))

}

//harus ditentuin lengthnya