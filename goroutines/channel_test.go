package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string) //untuk membuat channel
	defer close(channel) //untuk menutup channel

	// channel <- "fahmi" //untuk mengirim data ke channel
	// message := <- channel //untuk menerima data dari channel

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "fahmi" //harus ada yg menerimanya
		fmt.Println("selesai mengirim data")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string)  { // GiveMeResponse(paramenter type=chan type=string/int/dll)
	time.Sleep(2 * time.Second)
	channel <- "fahmi"
}

func TestChannelAsParameter(t *testing.T){
channel := make(chan string)
defer close(channel)

go GiveMeResponse(channel)

data := <-channel
fmt.Println(data)

time.Sleep(5 * time.Second)
}
// channel tidak butuh pointer karena defaultnya sudah by reference
//menggunakan channel jika menggunakan goroutine untuk menerima data
//harus ada yg ngirim ke channel dan harus ada yg menerima channel