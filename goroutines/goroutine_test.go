package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	println("Hello, World!")
}

// func TestCreateGoroutine(t *testing.T){
// 	go RunHelloWorld() // gunakan go di awal untuk asynchronous
// 	fmt.Println("Ups") // tidak akan menunggu func diatas selesai

// 	time.Sleep(1 * time.Second)
// }

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T){
	for i:= 1; i <= 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

//tidak cocok jika func mengugnakan goroutine mengembalikan value