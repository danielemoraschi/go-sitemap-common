package crawler

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/parser"
)

var mu = &sync.Mutex{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(urlsList *[]http.HttpResource, wg *sync.WaitGroup, url string, depth int,
    fetcher fetcher.FetcherInterface, parser parser.ParserInterface, policies []policy.PolicyInterface) {

    defer wg.Done()

    policy.UrlAllowedByPolicies(policies, url)

    fmt.Printf("Visiting: %s\n", url)

    urlResource, err := http.HttpResourceFactory(url, "")
    if err != nil {
        fmt.Println("HttpResource URL error: ", err)
        return
    }

    _, err = fetcher.Fetch(&urlResource)
    if err != nil {
        fmt.Println("Fetch error: ", err)
        return
    }

    urls, err := parser.Parse(urlResource)
    if err != nil {
        fmt.Println("Parse error: ", err)
        return
    }

    //fmt.Printf("Found: %s\n", urls)

    mu.Lock()
    *urlsList = append(*urlsList, urls...)
    mu.Unlock()

    if depth <= 0 {
        return
    }

    for _, u := range urls {
        if policy.UrlAllowedByPolicies(policies, u.String()) {
            wg.Add(1)
            go Crawl(urlsList, wg, u.String(), depth-1, fetcher, parser, policies)
        }
    }

    return
}
