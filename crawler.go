package crawler

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
)

var mu = &sync.Mutex{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(urlsList *[]string, wg *sync.WaitGroup, url string, depth int,
    fetcher fetcher.FetcherInterface, policies []policy.PolicyInterface) {

    defer wg.Done()

    policy.UrlAllowedByPolicies(policies, url)

    fmt.Printf("Visiting: %s\n", url)

    body, urls, err := fetcher.Fetch(http.HttpResourceFactory(url))
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Found: %s %q\n", urls, body)

    mu.Lock()
    *urlsList = append(*urlsList, urls...)
    mu.Unlock()

    if depth <= 0 {
        return
    }

    for _, u := range urls {
        if policy.UrlAllowedByPolicies(policies, u) {
            wg.Add(1)
            go Crawl(urlsList, wg, u, depth-1, fetcher, policies)
        }
    }

    return
}
