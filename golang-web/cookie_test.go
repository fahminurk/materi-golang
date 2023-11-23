package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// http.SetCookie(writer, &http.Cookie{
	// 	Name:  "token",
	// 	Value: "fahmi",
	// 	Path:  "/",
	// })

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "fahmi"
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success set cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("token")

	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "token: %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T){
req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
recorder := httptest.NewRecorder()

SetCookie(recorder, req)

cookies := recorder.Result().Cookies()

for _, cookie := range cookies {
	fmt.Printf("Cookie %s: %s \n", cookie.Name, cookie.Value)
}
}

func TestGetCookie(t *testing.T){
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "fahmi"
	req.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, req)

	body,_ := io.ReadAll(recorder.Body)

	fmt.Println(string(body))

}