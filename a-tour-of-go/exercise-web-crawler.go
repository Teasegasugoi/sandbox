package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	mu sync.Mutex
	c  map[string]bool
}

func (sc *SafeCache) isSame(url string) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	if sc.c[url] == true {
		return true
	}
	sc.c[url] = true
	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	sc := SafeCache{c: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	crawl(url, depth, fetcher, &sc, &wg)
	wg.Wait()
	return
}

func crawl(url string, depth int, fetcher Fetcher, sc *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// 既に取得したURLかどうか
	if sc.isSame(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	newWg := &sync.WaitGroup{}
	newWg.Add(len(urls))

	for _, u := range urls {
		go crawl(u, depth-1, fetcher, sc, newWg)
	}

	newWg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
