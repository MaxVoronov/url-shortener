package urlShortener

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
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
	normalizedUrl := normalizeUrl(url)

	hasher := md5.New()
	io.WriteString(hasher, normalizedUrl)
	hash := hex.EncodeToString(hasher.Sum(nil))[:7]

	s.storage[hash] = normalizedUrl
	return hash
}

func (s *UrlShortener) Resolve(url string) string {
	url = normalizeUrl(url)
	if result, ok := s.storage[url]; ok {
		return result
	}

	return ""
}

func normalizeUrl(url string) string {
	url = strings.ToLower(url)
	url = strings.TrimSpace(url)

	return url
}
