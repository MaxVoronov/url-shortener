package urlShortener

import "testing"

const rawUrl = "https://google.com/"

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

	if result := service.Shorten(rawUrl); result == "" {
		t.Errorf("Incorrect URL key: it can't be empty")
	}
}

func TestDuplicatedUrls(t *testing.T) {
	service := New()
	urlKey1 := service.Shorten(rawUrl)
	urlKey2 := service.Shorten(rawUrl)

	if urlKey1 == urlKey2 {
		t.Errorf("Duplicated URL keys: they must be different")
	}
}

func TestResolveSuccess(t *testing.T) {
	service := New()
	shortUrl := service.Shorten(rawUrl)

	if result := service.Resolve(shortUrl); result != rawUrl {
		t.Errorf("Incorrect resolved URL: expected %s, got %s", rawUrl, result)
	}
}
