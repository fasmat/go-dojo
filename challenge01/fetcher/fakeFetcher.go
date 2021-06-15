package fetcher

import "fmt"

// FakeFetcher is Fetcher that returns canned results.
type FakeFetcher struct {
	data map[string]*Result
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f.data[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func Fake() Fetcher {
	return FakeFetcher{data: data}
}
