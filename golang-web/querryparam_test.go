package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(writer, "hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleQueryParam(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	age := request.URL.Query().Get("age")

	fmt.Fprintf(writer, "Hello %s, you are %s", name, age)
}

func TestMultipleQueryParam(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko&age=20", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func MultipleParameterValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"] //mendapatkan semua query yg memiliki key name

	fmt.Fprint(writer, strings.Join(names, " "))

}

func TestMultipleParameterValue(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko&name=Kurniawan&name=Setiawan", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, req)

		res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}