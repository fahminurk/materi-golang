package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func startApp(error bool) {
	defer endApp() // defer tetap akan dijalankan jika terjadi error

	if error {
		panic("Ups")
	}

	fmt.Println("Start App")

}

func main(){
startApp(false)
startApp(true)
}