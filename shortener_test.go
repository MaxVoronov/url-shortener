package urlShortener

import "testing"

const rawUrl = "https://google.com/"
const hashedUrl = "f82438a"

func TestUrlNormalizing(t *testing.T) {
	if result := normalizeUrl(rawUrl); result != rawUrl {
		t.Errorf("Incorrect URL normalizing: expected %s, got %s", rawUrl, result)
	}

	unnormal := "  https://GOOGLE.com/   \t\n"
	if result := normalizeUrl(unnormal); result == unnormal {
		t.Errorf("Incorrect URL normalizing: expected %s, got %s", unnormal, result)
	}
}

func TestShortenSuccess(t *testing.T) {
	service := New()

	if result := service.Shorten(rawUrl); result != hashedUrl {
		t.Errorf("Incorrect hash: expected %s, got %s", hashedUrl, result)
	}
}

func TestResolveSuccess(t *testing.T) {
	service := New()
	shortUrl := service.Shorten(rawUrl)

	if result := service.Resolve(shortUrl); result != rawUrl {
		t.Errorf("Incorrect resolved URL: expected %s, got %s", rawUrl, result)
	}
}
