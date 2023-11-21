package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("salah")

	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println(err.Error())
	}

	binary := strconv.FormatInt(42, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(42)
	fmt.Println(stringInt)
}