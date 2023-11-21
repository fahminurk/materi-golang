package goroutines

import (
	"fmt"
	"testing"
)


func TestSelectChannel(t *testing.T){
channel1 := make(chan string)
channel2 := make(chan string)

go GiveMeResponse(channel1)
go GiveMeResponse(channel2)

counter := 0
for{
	select{
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
	}
	if counter == 2 {
		break
	}
}
}