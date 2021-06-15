package fetcher

import "fmt"

type DistinctFetcher struct {
	data    map[string]*Result
	fetched map[string]bool
}

func (f DistinctFetcher) Fetch(url string) (string, []string, error) {
	if f.fetched[url] {
		panic(fmt.Errorf("fetched same URL twice: %s", url))
	}

	f.fetched[url] = true

	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func (f DistinctFetcher) Completed() bool {
	for k := range data {
		if !f.fetched[k] {
			return false
		}
	}

	return true
}

func Distinct() DistinctFetcher {
	return DistinctFetcher{
		data:    data,
		fetched: make(map[string]bool),
	}
}
