package challenge01

import (
	"testing"
	"time"

	"dojo/challenge01/fetcher"

	"github.com/stretchr/testify/assert"
)

// Safety check. If this test fails you broke something in the Crawler.
func Test_Challenge01_00_Basic_Functionality(t *testing.T) {
	f := fetcher.Fake()

	c := New()
	c.Crawl("https://golang.org/", 4, f)
}

// The fetcher used in this test will panic when the same URL is fetched twice.
// Extend the Crawler, so that it fetches every URL only once.
func Test_Challenge01_01_Do_not_Fetch_URLs_Twice(t *testing.T) {
	f := fetcher.Distinct()

	c := New()
	c.Crawl("https://golang.org/", 4, f)

	assert.True(t, f.Completed(), "Not all URLs fetched")
}

// The fetcher used in this test also panics when an URL is fetched twice and additionally
// is very slow. Use concurrency to speed up the crawler and fetch multiple URLs simultaneously.
// The test fails after a timeout is reached to check if you are able to speed up the fetching process.
func Test_Challenge01_02_Be_More_Efficient(t *testing.T) {
	f := fetcher.Slow()

	done := make(chan bool)
	go func() {
		c := New()
		c.Crawl("https://golang.org/", 4, f)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		assert.Fail(t, "too slow!")
	}

	assert.True(t, f.Completed(), "Not all URLs fetched")
}

// The fetcher in this test simulates a server with a tight rate limit from a client, if you fetch more often
// than once per second it will fail. Introduce a rate limiter in the crawler that prevents this from happening
// without losing the functionality in the
func Test_Challenge01_03_RateLimit_Requests(t *testing.T) {
	f := fetcher.RateLimited()

	c := NewWithRateLimit(2 * time.Second)
	c.Crawl("https://golang.org/", 4, f)

	assert.True(t, f.Completed(), "Not all URLs fetched")
}
