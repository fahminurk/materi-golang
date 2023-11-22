/*
pool adalah implementasi design pattern bernama object pool pattern
pool untuk menyimpan data, bisa ambil dan gunakan datanya, setelah selesai kembalikan lgi ke poolnya
pool biasanya dipakai untuk koneksi ke database
pada pool tidak akan terjadi race condition
*/

package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T){
	// pool := sync.Pool{}

	//jika data di pool habis, bisa membuat defaultvalue =>
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("message 1") // put = memasukan data ke pool
	pool.Put("message 2")
	pool.Put("message 3")

	for i:=0; i < 10 ; i++ {
		go func ()  {
			data := pool.Get() // get = mengambil data dari pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // setelah selesai wajib put kembali
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}