package challenge01

import (
	"fmt"

	"dojo/challenge01/fetcher"
)

type Crawler01 struct {
	fetched map[string]bool
}

func New01() *Crawler01 {
	return &Crawler01{
		fetched: make(map[string]bool),
	}
}

func NewWithRateLimit01(...interface{}) Crawler {
	// Needed for TODO3
	return 0
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler01) Crawl(url string, depth int, fetcher fetcher.Fetcher) {
	if depth <= 0 {
		return
	}
	if c.checkFetched(url) {
		fmt.Printf("skipping: %s\n", url)
		return
	}
	c.fetched[url] = true

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

func (c *Crawler01) checkFetched(url string) bool {
	if c.fetched[url] {
		return true
	}
	c.fetched[url] = true
	return false
}
