package scrapers

import "net/http"


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
