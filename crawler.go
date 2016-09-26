package crawler

import (
    "fmt"
    "sync"
    "github.com/danielemoraschi/go-sitemap-common/policy"
    "github.com/danielemoraschi/go-sitemap-common/http"
    "github.com/danielemoraschi/go-sitemap-common/http/fetcher"
    "github.com/danielemoraschi/go-sitemap-common/parser"
)


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(semaphore chan bool, urlsList *UrlCollection, wg *sync.WaitGroup, url string, depth int,
    fetcher fetcher.FetcherInterface, parser parser.ParserInterface, policies []policy.PolicyInterface) {

    defer func() { <-semaphore }()
    defer wg.Done()

    policy.UrlAllowedByPolicies(policies, url)

    if depth <= 0 {
        return
    }

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

    urlsList.AddList(urls)

    for _, u := range urls {
        if policy.UrlAllowedByPolicies(policies, u.String()) {
            wg.Add(1)
            go Crawl(semaphore, urlsList, wg, u.String(), depth-1, fetcher, parser, policies)
        }
    }

    return
}
