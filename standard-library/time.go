package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Println(now)

	// var utc time.Time = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// fmt.Println(utc)

	formatter := "2006-01-02 15:04:05"
	value := "2023-01-01 00:00:00"
	valueTime, err := time.Parse(formatter,value)

	if err == nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println("error: ", err.Error())
	}
}