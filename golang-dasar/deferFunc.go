package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(){
	defer logging() // akan dijalankan setelah runApp() selesai, mau itu ada error ataupun tidak
	fmt.Println("runn aplikasi")
}

func main() {
runApp()
}