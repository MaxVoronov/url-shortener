# Otus Homework #1.2
## URL Shortener (модуль)

**Цель:** Разобраться, что такое и как работают интерфейсы, методы и типы. Впоследствии мы реализуем полноценный API
для сервиса сокращения ссылок. Это творческая задачка, которая потребует изучения документации, самостоятельной
работы и общения с преподавателями. Написать тип, который реализует интерфейс:

```go
type Shortener interface {
    Shorten(url string) string
    Resolve(url string) string
}
```

Метод _Shorten_ - возвращать "короткую" ссылку (выбор алгоритма - за студентом), например 
`otus.ru/some-long-link -> otus.ru/jhg34` и сохранять соответствие короткой и исходной ссылок в памяти
(не используя БД, а использовать, например, map). При вызове метода _Resolve_ - отдавать "длинную ссылку"
или пустую строку, если ссылка не найдена.

### Пример использования
```go
import "github.com/maxvoronov/url-shortener"

func main() {
	service := urlShortener.New()

	urlKey := service.Shorten("https://google.com/")
	if urlKey == "" {
		log.Panic("Failed to get URL key")
	}
	log.Printf("Short URL key: %s", urlKey)

	fullUrl := service.Resolve(urlKey)
	if fullUrl == "" {
		log.Panicf("Failed to get full URL by key %s", urlKey)
	}
	log.Printf("Full URL: %s", fullUrl)
}

// Result output:
// 2019/06/05 22:39:45 Short URL key: f82438a
// 2019/06/05 22:39:45 Full URL: https://google.com/
```
