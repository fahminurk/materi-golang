package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//atomic cocok untuk type data primitif
// kalau struct dan dll gunakan mutex

func TestAtomic(t *testing.T){
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func ()  {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				// x += 1 // terkena race condition
				atomic.AddInt64(&x, 1)
			}	
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("counter: ", x)
}