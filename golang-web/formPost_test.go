package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	//cara manual
	name := request.PostForm.Get("name")
	password := request.PostForm.Get("password")

	fmt.Fprintf(writer, "Hello %s, your password is %s", name, password)

}

func TestFormPost(t *testing.T){
	requestBody := strings.NewReader("name=joko&password=secret")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, req)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}