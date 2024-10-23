package main

import (
	"fmt"
	"net/http"

	"github.com/taiypeo/url-shortener/storage"
)

const ID_PATHVALUE = "id"

func main() {
	urlStorage := storage.NewLocalStorage()
	http.HandleFunc(
		"POST /{$}",
		func(w http.ResponseWriter, req *http.Request) {
			createURL(urlStorage, w, req)
		},
	)
	http.HandleFunc(
		fmt.Sprintf("GET /{%s}", ID_PATHVALUE),
		func(w http.ResponseWriter, req *http.Request) {
			redirectURL(urlStorage, w, req)
		},
	)
	http.ListenAndServe(":8080", nil)
}
