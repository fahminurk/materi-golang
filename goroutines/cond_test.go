/*
cond adalah implementasi locking berbasis kondisi

*/

package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var group = sync.WaitGroup{}
var cond = sync.NewCond(&sync.Mutex{})

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)
	cond.L.Lock() //kondisi di lock
	cond.Wait() // kondisi di tunggu

	fmt.Println("done", value)

	cond.L.Unlock() //kondisi di unlock
}

func TestCond(t *testing.T){

	for i:=0; i < 10; i++{
		go WaitCondition(i)
	}

	// go func ()  {
	// 	for i:=0; i < 10; i++{
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal() //memberikan signal agar 1 goroutine lanjut 
	// 	}	
	// }()

	go func ()  {
	time.Sleep(1 * time.Second)
	cond.Broadcast() //memberikan signal agar semua goroutine lanjut	
	}()

	group.Wait()
}