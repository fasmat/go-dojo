package fetcher

import (
	"fmt"
	"sync"
	"time"
)

type RateLimitedFetcher struct {
	data    map[string]*Result
	fetched sync.Map

	mu        sync.Mutex
	lastFetch time.Time
}

func (f *RateLimitedFetcher) Fetch(url string) (string, []string, error) {
	if _, ok := f.fetched.Load(url); ok {
		panic(fmt.Errorf("fetched same URL twice: %s", url))
	}
	f.fetched.Store(url, true)

	t := f.updateLastFetched()
	if time.Since(t).Milliseconds() < 1000 {
		panic(fmt.Errorf("fetch rate too high"))
	}

	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func (f *RateLimitedFetcher) Completed() bool {
	for k := range data {
		if _, ok := f.fetched.Load(k); !ok {
			return false
		}
	}

	return true
}

func (f *RateLimitedFetcher) updateLastFetched() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()

	t := f.lastFetch
	f.lastFetch = time.Now()
	return t
}

func RateLimited() *RateLimitedFetcher {
	return &RateLimitedFetcher{
		data: data,
	}
}
