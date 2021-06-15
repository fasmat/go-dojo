package challenge01

import (
	"fmt"

	"dojo/challenge01/fetcher"
)

type Crawler int

func New() Crawler {
	return 0
}

func NewWithRateLimit(...interface{}) Crawler {
	// Needed for TODO3
	return 0
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c Crawler) Crawl(url string, depth int, fetcher fetcher.Fetcher) {
	// TODO1: Don't fetch the same URL twice.
	// TODO2: Fetch URLs in parallel.
	// TODO3: Add an option to rate limit request to the server.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		c.Crawl(u, depth-1, fetcher)
	}
}
