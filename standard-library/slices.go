package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Jane", "Bob"}
	numbers := []int{1, 2, 3, 10}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Contains(names, "John"))
	fmt.Println(slices.Index(names, "John"))
}