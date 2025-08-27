package scrapers

import (
	"image"
)

type Book struct {
	Id string
	Slug string
	Title string
}

type Chapter struct {
	Id string
	Info string
	Language string
	Number string
}

type ScraperFuncs struct {
	Search func(string) []Book
	Chapters func(string) []Chapter
	Pages func(string, string)
	Interactive func(string) []image.Image
}

var Registry = make(map[string]ScraperFuncs)

func Register(name string, fns ScraperFuncs) {
	Registry[name] = fns
}
