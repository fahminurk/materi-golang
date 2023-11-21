package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "fahmi,nurkamil\n" +
		"john,doe\n" +
		"jane,doe\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for{
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	
}