package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(3.5)) //bulatkan terdekat
	fmt.Println(math.Ceil(3.6)) //bulatkan keatas
	fmt.Println(math.Floor(3.4)) //bulatkan kebawah
	fmt.Println(math.Max(1, 2)) //mencari nilai terbesar
	fmt.Println(math.Min(1, 2)) //mencari nilai terkecil
	fmt.Println(math.Pow(2, 3)) //pangkat
}