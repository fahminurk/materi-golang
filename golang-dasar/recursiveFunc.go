package main

import "fmt"

func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

func lops(n int) int {
	result := 1

	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	factorial := recursive(5)
	fmt.Println(factorial)
	result := lops(5)
	fmt.Println(result)
}