package belaejar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(body)
}
func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000?name=ucup", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(body)
}
