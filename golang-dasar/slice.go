package main

import (
	"fmt"
)

func main() {
	names := [...]string{"fahmi", "nurkamil", "udin", "petot", "budi"}

	slice1 := names[0:4]
	fmt.Println(slice1)
	slice2 := names[4:5]
	fmt.Println(slice2)
	slice3 := names[3:]
	fmt.Println(slice3)
	var slice4 []string = names[:]
	fmt.Println(slice4)

	// function slice
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:] // [sabtu, minggu]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"

	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru") //karena capacitynya melebihi, jadi membuat array baru tapi tidak merubah araay utamanya
	fmt.Println(daySlice2)

	//membuat slice dari awal
	var newSlice = make([]string, 2, 5) // make([]typedata, length, capacity)
	newSlice[0] = "fahmi"
	newSlice[1] = "nurkamil"
	// newSlice[2] = "udin" //tidak bisa karena length nya 2, harusn menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "udin")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//perbedaan membuat array dan slice, hati2
	isArray := [...]int{1, 2, 3, 4, 5} // harus di isi lengthnya untk array
	isSLice := []int{1, 2, 3, 4, 5}

	fmt.Println(isArray)
	fmt.Println(isSLice)
}

// pointer = penunjuk data pertama di array
// length = panjan dari slicenya
// capacity = kapasitas dari slicenya, dimana length
// 			tidak boleh lebih dari capacity