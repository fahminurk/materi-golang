package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	// data.Value = 1

	// data = data.Next()
	// data.Value = 2

	// data = data.Next()
	// data.Value = 3

	// data = data.Next()
	// data.Value = 4

	// data = data.Next()
	// data.Value = 5

	
	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.Itoa(i)
		data = data.Next()
	}
	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}