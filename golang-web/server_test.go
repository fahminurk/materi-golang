package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
server := http.Server{
	Addr: "localhost:8080",
	 
}

err := server.ListenAndServe()

if err != nil {
	panic(err)
}
}

func TestHandler(t *testing.T){

	var handler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World!")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T){

	var handler http.HandlerFunc = func(writer http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(writer, r.Method) //get post put delete
		fmt.Fprint(writer, r.RequestURI) // /users /images
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}