package main

import "fmt"

//tidak bisa return string
// func nama(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	data := newMap("fahmi")
	fmt.Println(data)
}

// nil hanya bisa dipakai di interface, func, map, slice, pointer, channel