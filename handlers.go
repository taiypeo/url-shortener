package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func createURL(w http.ResponseWriter, req *http.Request) {
	const maxShorteningAttempts = 5

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

	mut.Lock()
	defer mut.Unlock()
	for range maxShorteningAttempts {
		shortURL := buildShortenedURL()
		if _, ok := urls[shortURL]; !ok {
			urls[shortURL] = body

			log.Printf("POST 201 / \"%s\"", body)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "%s\n", shortURL)
			return
		}
	}

	log.Printf("POST 500 / \"%s\"", body)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%v!\n", err)
}

func redirectURL(w http.ResponseWriter, req *http.Request) {
	shortURL := req.PathValue(ID_PATHVALUE)

	mut.Lock()
	defer mut.Unlock()
	if url, ok := urls[shortURL]; ok {
		log.Printf("GET 301 /%s", shortURL)
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusMovedPermanently)
	} else {
		log.Printf("GET 404 /%s", shortURL)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found")
	}
}
