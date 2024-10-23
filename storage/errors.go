package storage

import "fmt"

var (
	ErrExceededMaxShorteningAttempts = fmt.Errorf("unable to insert URL, exceeded maxShorteningAttempts")
	ErrShortURLNotFound              = fmt.Errorf("short URL not found")
)
