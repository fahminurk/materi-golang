package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "asuudahlah,a,sdd,cx"

	fmt.Println(strings.Contains(str, "asu"))
	fmt.Println(strings.Count(str, "a"))
	fmt.Println(strings.Split(str, ","))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Trim(str, "u"))
	fmt.Println(strings.ReplaceAll(str, "u", "o"))
}