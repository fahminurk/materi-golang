package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/notfound.html
var notfoundHTML string

func ServeFile(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		fmt.Fprint(writer, notfoundHTML)
	}
}

func TestServeFileServer(t *testing.T){

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}