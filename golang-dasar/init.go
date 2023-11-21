package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal" //gunakan _ (blank identifier) untuk menonaktifkan internal
)

func main(){
	fmt.Println(database.GetDatabase())
}