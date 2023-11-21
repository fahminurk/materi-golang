package main

import "fmt"

//variable argumen, menerima banyak argumen tetapi harus 1 type data
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(slice...))
}