package scrapers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func search(query string) []Book {
	headers := BaseHeaders.Clone()
	headers.Set("Referer", "https://comick.io/")
	headers.Set("Origin", "https://comick.io")

	baseURL := "https://api.comick.io/v1.0/search"
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	formattedQuery := query

	q := u.Query()
	q.Set("q", formattedQuery)
	// i don't know what exactly it is, but if set to false, response has additional information
	q.Set("t", "true")

	u.RawQuery = q.Encode()


	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header = headers


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	type Data struct {
		Id string `json:"hid"`
		Title string `json:"title"`
		Type string `json:"type"`
	}

	var data []Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	books := make([]Book, 0, len(data))

	for _, d := range data {
		if d.Type == "author" {
			continue
		}
		books = append(books, Book{
			Id: d.Id,
			Title: d.Title,
		})

	}

	return books
}

// TODO: probably pages of these requests

func chapters(id string) []Chapter {
	headers := BaseHeaders.Clone()
	headers.Set("Referer", "https://comick.io/")
	headers.Set("Origin", "https://comick.io")

	baseURL := fmt.Sprintf("https://api.comick.io/comic/%s/chapters", id)
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}


	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header = headers


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// TODO: handle if id is wrong
	fmt.Print(string(body))


	type ChapterWrapper struct {
		Id string `json:"hid"`
		Title string `json:"chap"`
		Language string `json:"lang"`
		Groups []string `json:"group_name"`
	}

	type Data struct {
		Chapters []ChapterWrapper `json:"chapters"`
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	chaps := make([]Chapter, 0, len(data.Chapters))

	for _, c := range data.Chapters {
		chaps = append(chaps, Chapter{
			Id: c.Id,
			Title: c.Title,
			Info: fmt.Sprintf("%s, [%s]", c.Language, strings.Join(c.Groups, ", ")),
		})

	}

	return chaps




}

func init() {
	Register(
		"comick.io",
		ScraperFuncs{
			Search: search,
			Chapters: chapters,
		},
	)
}

