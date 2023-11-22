package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T){
	background := context.Background()
	// fmt.Println(background)

	// todo := context.TODO()
	// fmt.Println(todo)

	ctxA := context.WithValue(background, "a", "A")
	ctxB := context.WithValue(background, "b", "B")

	ctxA1 := context.WithValue(ctxA, "a1", "A1")
	ctxA2 := context.WithValue(ctxA, "a2", "A2")
	ctxB1 := context.WithValue(ctxB, "b1", "B1")

	fmt.Println(ctxA1.Value("a1")) //untuk menggambil data dari context
	fmt.Println(ctxA2.Value("a"))
	fmt.Println(ctxB1)

}

/*
saat menambahkan data ke context, maka context tersebuat akan membuat child baru
dan data parentnya tidak akan berubah, karena data di context bersifat imutable
jika ingin menggambil data, hanya bisa dari child dan parentnya
tidak bisa dari child ke parent
*/

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T){
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx,cancel := context.WithCancel(parent)// ngecancel context jika sudah selesai
	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T){
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx,cancel := context.WithTimeout(parent, 5*time.Second)// ngecancel context dengan timeout
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter", n)	
	}
	
	time.Sleep(1 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}
func TestContextWithDeadline(t *testing.T){
	fmt.Println("total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx,cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))// ngecancel context dengan timeout
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter", n)	
	}
	
	time.Sleep(1 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}