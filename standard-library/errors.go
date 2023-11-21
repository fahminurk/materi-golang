package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found")
)

func getById(id string) error {

	if id == "" {
		return ValidationError
		} else if id == "1" {
			return NotFoundError
	} else {
		return nil
	}
}

func main(){
err := getById("")

if err != nil {
	if errors.Is(err, ValidationError) {
		fmt.Println("validation error")
	} else if errors.Is(err, NotFoundError) {
		fmt.Println("not found")
	}
	
}
}