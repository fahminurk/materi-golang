package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) { //hanya berjalan di satu package atau package itu sendiri
	//before
fmt.Println("before unit test")
	m.Run()
	//after
	fmt.Println("after unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("cannot run on windows")
	}
	result := HelloWorld("fahmi")
	require.Equal(t, "Hello fahmi", result, "result must be 'Hello fahmi'")
	fmt.Println("Test Hello require Success!")
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("fahmi")
	require.Equal(t, "Hello fahmi", result, "result must be 'Hello fahmi'")
	fmt.Println("Test Hello require Success!")
}
func TestHelloAssert(t *testing.T) {
	result := HelloWorld("fahmi")
	assert.Equal(t, "Hello fahmi", result, "result must be 'Hello fahmi'")
	fmt.Println("Test Hello Assert Success!")
}

// func TestHelloWorld(t *testing.T) {
// result := HelloWorld("fahmi")

// if result != "Hello fahmi" {
// 	// t.Fail()
// 	t.Error("result must be 'Hello fahmi'")
// }
// }
// func TestHellobudi(t *testing.T) {
// result := HelloWorld("budi")

// if result != "Hello budii" {
// 	// t.FailNow()
// 	t.Fatal("result must be 'Hello budi'")
// }
// }

// cd helper => go test -v
// cd ... => go test -v ./... untuk di semua package/folder

//assertion
// go get github.com/stretchr/testify