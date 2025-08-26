package scrapers

import (
	"net/http"

	"image"
	_ "image/png"
	_ "image/jpeg"

	"sync"
)


var BaseHeaders = http.Header{}

func init() {

	// BaseHeaders.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:141.0) Gecko/20100101 Firefox/141.0")

	// chrome on windows for anti fingerprinting or whatever
	BaseHeaders.Set("User-Agent", "Mozilla/5.0 (Windows NT 11.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.6998.166 Safari/537.36")

	BaseHeaders.Set("Accept", "*/*")
	BaseHeaders.Set("Accept-Language", "en-US,en;q=0.5")
	BaseHeaders.Set("DNT", "1")
	BaseHeaders.Set("Sec-GPC", "1")
	BaseHeaders.Set("Sec-Fetch-Dest", "empty")
	BaseHeaders.Set("Sec-Fetch-Mode", "cors")
	BaseHeaders.Set("Sec-Fetch-Site", "same-site")
	BaseHeaders.Set("Connection", "keep-alive")
}


func getImage(urlString string, headers http.Header) (image.Image, error) {

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header = headers


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}


	return img, nil
}

var MaxConcurrencyPerImageSlice = 5


func getImages(urls []string, headers http.Header) ([]image.Image, error) {

	sem := make(chan struct{}, MaxConcurrencyPerImageSlice)

	var wg sync.WaitGroup

	images := make([]image.Image, len(urls))
	errors := make([]error, len(urls))

	var lastError error

	wg.Add(len(urls))
	for i, url := range urls {

		go func(i int, url string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <- sem }()

			img, err := getImage(url, headers)
			if err != nil {
				errors[i] = err
				lastError = err
				return
			}

			images[i] = img


		}(i, url)

	}

	wg.Wait()



	return images, lastError

}
