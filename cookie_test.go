package belaejar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-YUSUF-NAME"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-YUSUF-NAME")
	if err != nil {
		fmt.Fprint(w, "no cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/?name=ucup", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "X-YUSUF-NAME"
	cookie.Value = "yusuf"
	request.AddCookie(cookie)

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Print(string(body))
}