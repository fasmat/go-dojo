package challenge01

import (
	"fmt"
	"sync"
	"time"

	"dojo/challenge01/fetcher"
)

type Crawler03 struct {
	fetched sync.Map
	wg      sync.WaitGroup
	ticker  *time.Ticker
}

func New03() *Crawler03 {
	c := &Crawler03{}
	return c
}

func NewWithRateLimit03(rate time.Duration) *Crawler03 {
	c := &Crawler03{
		ticker:  time.NewTicker(rate),
	}
	return c
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler03) Crawl(url string, depth int, fetcher fetcher.Fetcher) {
	c.wg.Add(1)
	go c.crawlHandler(url, depth, fetcher)
	c.wg.Wait()
}

func (c *Crawler03) crawlHandler(url string, depth int, fetcher fetcher.Fetcher) {
	defer c.wg.Done()
	if depth <= 0 {
		return
	}

	if c.fetchedBefore(url) {
		fmt.Printf("skipping: %s\n", url)
		return
	}

	if c.ticker != nil {
		<-c.ticker.C
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	c.wg.Add(len(urls))
	for _, u := range urls {
		go c.crawlHandler(u, depth-1, fetcher)
	}
}

func (c *Crawler03) fetchedBefore(url string) bool {
	if _, ok := c.fetched.Load(url); ok {
		return true
	}
	c.fetched.Store(url, true)
	return false
}
