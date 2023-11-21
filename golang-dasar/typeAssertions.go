package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultAsString := result.(string)
	fmt.Println(resultAsString)

	// resultAsInt := result.(int) //akan terjadi panic 
	// fmt.Println(resultAsInt)

	//menggunakan switch untuk safety agar tidak terjadi panic
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}