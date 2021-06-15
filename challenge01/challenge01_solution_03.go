package challenge01

import (
	"fmt"
	"sync"
	"time"

	"dojo/challenge01/fetcher"
)

type Crawler03 struct {
	fetched map[string]bool
	wg      *sync.WaitGroup
	mu      sync.Mutex
	ticker  *time.Ticker
}

func New03() *Crawler03 {
	c := &Crawler03{
		fetched: make(map[string]bool),
		wg:      new(sync.WaitGroup),
	}
	c.wg.Add(1)
	return c
}

func NewWithRateLimit03(rate time.Duration) *Crawler03 {
	c := &Crawler03{
		fetched: make(map[string]bool),
		wg:      new(sync.WaitGroup),
		ticker:  time.NewTicker(rate),
	}
	c.wg.Add(1)
	return c
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler03) Crawl(url string, depth int, fetcher fetcher.Fetcher) {
	go c.crawHandler(url, depth, fetcher)

	c.wg.Wait()
}

func (c *Crawler03) crawHandler(url string, depth int, fetcher fetcher.Fetcher) {
	defer c.wg.Done()
	if depth <= 0 {
		return
	}

	if c.checkFetched(url) {
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
		go c.Crawl(u, depth-1, fetcher)
	}
}

func (c *Crawler03) checkFetched(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.fetched[url] {
		return true
	}
	c.fetched[url] = true
	return false
}
