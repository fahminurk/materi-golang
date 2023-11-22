package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T){
	once := sync.Once{} //sebuah func hanya di eksekusi sekali, go yg pertama yg bisa mengaksesnya
	group := sync.WaitGroup{}

	for i := 0; i <100; i++ {
		go func( ){
			group.Add(1)
			once.Do(OnlyOnce) //hanya func yg tidak memiliki parameter
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("counter :",counter)
}