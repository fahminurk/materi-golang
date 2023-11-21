package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{
			Message: "ID cannot be empty",
		}
	}

	if id != "fahmi" {
		return &notFoundError{"ID not found"}
	}

	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	println(validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	println(notFoundErr.Message)
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			println(finalError.Error())
		case *notFoundError:
			println(finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	}
}