package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
	fmt.Println("selesai")

	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan ke", i)
	}

	names := []string{"fahmi", "nur", "kamil"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println(name)
	}

	//break & continue
	for i:=0; i<10; i++ {
		if i == 5 {break}
		fmt.Println("perulangan ke", i)
	}

	for i:=0; i<10; i++ {
		if i%2 == 0 {continue}
		fmt.Println("perulangan ke", i)
	}
}