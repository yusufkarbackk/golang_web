package belaejar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "hello world")
}

func TestHttp(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/menu", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}