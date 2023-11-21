package main

import "fmt"

type Man struct {
	Name string
}
type Girl struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (girl *Girl) Married(){ //direkomen selalu menggunakan pointer
girl.Name = "Miss. " + girl.Name
}

func main() {
	fahmi := Man{"fahmi"}
	fahmi.Married()
	fmt.Println(fahmi.Name) // => fahmi

	ukhti := Girl{"ukhti"}
	ukhti.Married()
	fmt.Println(ukhti.Name) // => miss. ukhti

}