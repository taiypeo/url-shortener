package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/taiypeo/url-shortener/storage"
)

func createURL(urlStorage storage.Storage, w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("POST 500 /")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v!\n", err)
		return
	}

	body := string(bytes)
	if !strings.HasPrefix(body, "http") {
		body = "http://" + body
	}

	if shortUrl, err := urlStorage.CreateShortURL(req.Context(), body); err != nil {
		log.Printf("POST 500 / \"%s\"", body)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v!\n", err)
	} else {
		log.Printf("POST 201 / \"%s\"", body)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s\n", shortUrl)
	}
}

func redirectURL(urlStorage storage.Storage, w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	shortUrl := req.PathValue(ID_PATHVALUE)
	if fullUrl, err := urlStorage.GetFullURL(req.Context(), shortUrl); err != nil {
		if errors.Is(err, storage.ErrShortURLNotFound) {
			log.Printf("GET 404 /%s", shortUrl)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Not Found")
		} else {
			log.Printf("POST 500 /%s", shortUrl)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v!\n", err)
		}
	} else {
		log.Printf("GET 301 /%s", shortUrl)
		w.Header().Set("Location", fullUrl)
		w.WriteHeader(http.StatusMovedPermanently)
	}
}
