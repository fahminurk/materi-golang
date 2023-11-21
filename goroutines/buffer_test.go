// secara default channel hanya bisa menerima 1 data, sisanya ngantri
// buffer digunakan untuk menampung data antrian di channel
package goroutines

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3) //bisa menerima 3 data tegantung buffer yg di isi
	defer close(channel)

	channel <- "fahmi"
	channel <- "budi"
	channel <- "andi"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("slesai")
}