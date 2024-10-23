package main

import (
	"fmt"
	"net/http"

	"github.com/taiypeo/url-shortener/storage"
)

const ID_PATHVALUE = "id"

var urlStorage storage.Storage

func main() {
	urlStorage = storage.NewLocalStorage()

	http.HandleFunc("POST /{$}", createURL)
	http.HandleFunc(fmt.Sprintf("GET /{%s}", ID_PATHVALUE), redirectURL)
	http.ListenAndServe(":8080", nil)
}
