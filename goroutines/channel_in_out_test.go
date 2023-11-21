package goroutines

import (
	"testing"
	"time"
)

//mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "fahmi"
}

//menggambil data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	println(data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}