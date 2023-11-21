package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronpus(group *sync.WaitGroup){
	defer group.Done()

	group.Add(1)

	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronpus(&group)
	}

	group.Wait() //akan menunggu sampai semua go selesai (tidak perlu menggunakan time.sleep)
	fmt.Println("done")
}