package urlShortener

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type UrlShortener struct {
	storage map[string]string
}

func New() Shortener {
	service := new(UrlShortener)
	service.storage = make(map[string]string)

	return service
}

func (s *UrlShortener) Shorten(url string) string {
	urlKey := ""
	isUsedKey := true
	normalizedUrl := normalizeUrl(url)

	for isUsedKey {
		urlKey = generateShortKey(normalizedUrl)
		_, isUsedKey = s.storage[urlKey]
	}
	s.storage[urlKey] = normalizedUrl

	return urlKey
}

func (s *UrlShortener) Resolve(url string) string {
	url = normalizeUrl(url)

	return s.storage[url]
}

func normalizeUrl(url string) string {
	url = strings.ToLower(url)
	url = strings.TrimSpace(url)

	return url
}

func generateShortKey(url string) string {
	hasher := md5.New()
	io.WriteString(hasher, url)
	io.WriteString(hasher, time.Now().String())
	hash := hex.EncodeToString(hasher.Sum(nil))[:7]

	return hash
}
