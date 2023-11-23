package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == ""{
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer,"name is empty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Joko", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
}