package fetcher

import (
	"fmt"
	"sync"
	"time"
)

type SlowFetcher struct {
	data    map[string]*Result
	fetched sync.Map
}

func (f *SlowFetcher) Fetch(url string) (string, []string, error) {
	if _, ok := f.fetched.Load(url); ok {
		panic(fmt.Errorf("fetched same URL twice: %s", url))
	}
	f.fetched.Store(url, true)

	<-time.After(1 * time.Second)

	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func (f *SlowFetcher) Completed() bool {
	for k := range data {
		if _, ok := f.fetched.Load(k); !ok {
			return false
		}
	}

	return true
}

func Slow() *SlowFetcher {
	return &SlowFetcher{
		data: data,
	}
}
