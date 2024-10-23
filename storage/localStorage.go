package storage

import (
	"context"
	"math/rand"
	"strings"
	"sync"
)

const maxShorteningAttempts = 5

func buildShortenedURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	const shortenedURLLength = 6

	sb := strings.Builder{}
	sb.Grow(shortenedURLLength)
	for range shortenedURLLength {
		sb.WriteByte(chars[rand.Intn(len(chars))])
	}

	return sb.String()
}

type LocalStorage struct {
	FullURLs map[string]string
	mut      sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{FullURLs: make(map[string]string)}
}

func (storage *LocalStorage) CreateShortURL(ctx context.Context, url string) (string, error) {
	storage.mut.Lock()
	defer storage.mut.Unlock()

	for range maxShorteningAttempts {
		shortURL := buildShortenedURL()
		if _, ok := storage.FullURLs[shortURL]; !ok {
			storage.FullURLs[shortURL] = url
			return shortURL, nil
		}
	}

	return "", ErrExceededMaxShorteningAttempts
}

func (storage *LocalStorage) GetFullURL(ctx context.Context, shortUrl string) (string, error) {
	storage.mut.RLock()
	defer storage.mut.RUnlock()

	if fullUrl, ok := storage.FullURLs[shortUrl]; ok {
		return fullUrl, nil
	}

	return "", ErrShortURLNotFound
}
