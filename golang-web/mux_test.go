package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// multiple endpoint
func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "welcome to my API!!")
	})

	mux.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "get all users")
	})

	// jika ada " / " di akhir, maka endpoint yg tidak ada akn ke images
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "images")
	})
	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}