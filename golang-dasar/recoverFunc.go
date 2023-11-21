package main

import "fmt"

func endAppp() {
	fmt.Println("End App")
	//cara benar
	message := recover()
	fmt.Println("terjadi panic", message)
}

func startAppp(error bool) {
	defer endAppp() // defer tetap akan dijalankan jika terjadi error

	if error {
		panic("Ups")
	}
	//cara salah
	// message := recover()
	// fmt.Println("terjadi panic", message)

}

func main(){
// startAppp(false)
startAppp(true)
fmt.Println("program tetap berjalan")
}

//defer recover panic sama kayak trycath di js