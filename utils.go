package main

import (
	"math/rand"
	"strings"
)

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
