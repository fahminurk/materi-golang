package main

import "fmt" 

func sayHello(name string)  {
	 fmt.Println("hello", name)
}

func getHello(name string) string {
	return "hello " + name
}
// multiple return
func getUser()(string, int){
	return "fahmi", 21
}

// func getFullName()(firstname string, lastname string, username string){
// 	firstname = "fahmi"
// 	lastname = "nurkamil"
// 	username = "fahminurkamil"
// 	return firstname, lastname, username
// }
func getFullName()(firstname , lastname , username string){
	firstname = "fahmi"
	lastname = "nurkamil"
	username = "fahminurk"
	return firstname, lastname, username
}

func main() {
	sayHello("budi")
	fmt.Println(getHello("fahmi"))

	name, age := getUser()
	fmt.Println(name, age)

	nama, _ := getUser() //hiraukan menggunakan _
	fmt.Println(nama)

	a,b,c := getFullName()
	fmt.Println(a, b, c)
}