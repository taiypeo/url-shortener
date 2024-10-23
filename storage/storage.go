package storage

import "context"

type Storage interface {
	CreateShortURL(ctx context.Context, url string) (string, error)
	GetFullURL(ctx context.Context, shortUrl string) (string, error)
}
