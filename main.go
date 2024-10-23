package main

import (
	"fmt"
	"net/http"
	"sync"
)

const ID_PATHVALUE = "id"

var urls map[string]string
var mut sync.Mutex

func main() {
	urls = make(map[string]string)

	http.HandleFunc("POST /{$}", createURL)
	http.HandleFunc(fmt.Sprintf("GET /{%s}", ID_PATHVALUE), redirectURL)
	http.ListenAndServe(":8080", nil)
}
