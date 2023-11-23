package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
fmt.Println(version)
}

//go:embed image.jpg
var image []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("image_new.jpg", image, fs.ModePerm)

	if err != nil {
		panic(err)
	}

}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries{
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}


/*
embed harus di luar func
embed = membaca sebuah file
go:embed version.txt harus di atas variable
*/