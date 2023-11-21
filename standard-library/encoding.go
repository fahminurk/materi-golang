package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "fahmi nurkamil"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	decoded, err := base64.StdEncoding.DecodeString(encoded)


	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(encoded)
	fmt.Println(string(decoded))
}